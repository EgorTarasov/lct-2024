from math import log
import os
import dotenv
import logging
from kafka import KafkaConsumer
from kafka.consumer.fetcher import ConsumerRecord
from typing import TypedDict
import datetime as dt
import json
import psycopg
from minio import Minio
import pandas as pd

from inference import ModelInference
from heat_features import load_heat_df, gen_features_heat_for_group
from address import load_address, process_address
from targets import load_targets, process_targets, load_processed_targets

dotenv.load_dotenv()


logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)
formatter = logging.Formatter("%(asctime)s - %(name)s - %(levelname)s - %(message)s")
ch = logging.StreamHandler()
ch.setFormatter(formatter)
logger.addHandler(ch)


class PredictionRequest(TypedDict):
    unoms: list[int]
    region_name: str
    prediction_id: int
    start_date: dt.date
    end_date: dt.date
    threshold: float


class UploadTableRequest(TypedDict):
    upload_id: int
    s3_key: str


def record_to_tuple(record: dict, upload_id: int):
    return (
        record["description"],
        record["system"],
        record["external_created_at"],
        record["completed"],
        record["region"],
        record["unom"],
        record["address"],
        record["external_completed"],
        upload_id,
    )


def process_events_table(file_path: str):
    column_names = {
        "Наименование": "description",
        "Источник": "system",
        "Дата создания во внешней системе": "external_created_at",
        "Дата закрытия": "completed",
        "Округ": "region",
        "УНОМ": "unom",
        "Адрес": "address",
        "Дата и время завершения события во внешней системе": "external_completed",
    }

    df = pd.read_excel(
        file_path,
        sheet_name="Выгрузка",
    )
    df = df.rename(columns=column_names)

    df["external_created_at"] = pd.to_datetime(df["external_created_at"])
    df["completed"] = pd.to_datetime(df["completed"])
    df["external_completed"] = pd.to_datetime(df["external_completed"])
    df["unom"] = df["unom"].dropna()
    df = df.dropna(subset=["unom"])
    df["unom"] = df["unom"].astype(int)

    return df.to_dict(orient="records")


# example table:
# ../data/dataset/5. Перечень событий за период 01.10.2023-30.04.2023 (ЦУ КГХ)/События за период_01.01.2024-30.04.2024.xlsx
EVENTS_TABLE_COLUMNS = [
    "Наименование",
    "Источник",
    "Дата создания во внешней системе",
    "Дата закрытия",
    "Округ",
    "УНОМ",
    "Адрес",
    "Дата и время завершения события во внешней системе",
]
EVENTS_TABLE_SHEET_NAMES = ["Сводная таблица", "Выгрузка", "Query"]

# example table
# ../data/dataset/11.Выгрузка_ОДПУ_отопление_ВАО_20240522.xlsx
HEAT_TABLE_SHEETS = ["Sheet 1", "Sheet 2", "Справочник Ошибки (W)"]

# example table
# ../data/dataset/13. Адресный реестр объектов недвижимости города Москвы.xlsx
ADDRESS_TABLE_SHEETS = ["0"]


def to_sql_record(record: dict):  # -> Generator[tuple[Any, Any, Any, Any], Any, None]:
    """
    insert into incidents(
    unom,
    opened_at,
    predicted_at,
    title,
    priority,
    status
    """
    for unom in record:
        date = record[unom]["date"]
        for event in record[unom]["events"]:
            value = record[unom]["events"][event]
            yield unom, date, event, value


class Controller:

    def __init__(
        self,
        consumer: KafkaConsumer,
        model: ModelInference,
        s3: Minio,
        dsn: str,
    ):
        self.consumer = consumer
        self.model = model
        self.dsn = dsn
        self.s3 = s3

    def __save_results(self, results: dict, region_name: str):
        table_name = "incident"

        query = """
insert into incident (description, system, external_created, region_name, unom, address, predicted_at)
values (%s, %s, %s, %s, %s, %s, %s);
"""

        with psycopg.connect(self.dsn) as conn:
            with conn.cursor() as cur:
                logger.info(f"uploading results")
                for unom, predicted_at, description, value in to_sql_record(results):
                    values = (
                        description,
                        "ml",
                        dt.datetime.now(),
                        region_name,
                        unom,
                        "ml",
                        dt.datetime.now(),
                    )
                    cur.execute(
                        query,
                        values,
                    )

            conn.commit()
            logger.info("saved results")

    def __parse_headers(self, headers: list | None) -> dict:
        headers_dict = {}
        if headers is None:
            return headers_dict
        for key, value in headers:
            headers_dict[key] = value.decode("utf-8")
        return headers_dict

    def __update_upload_status(self, upload_id: int, status: str):
        with psycopg.connect(self.dsn) as conn:

            with conn.cursor() as cur:
                cur.execute(
                    """update uploads
                            set status = %s
                            where id = %s;
                    """,
                    (
                        status,
                        upload_id,
                    ),
                )
            conn.commit()

    def __define_table_type(self, file_path: str) -> str:
        """Определения типа таблицы для дальнейшей обработки
        возможные типы таблиц:
        - target "data/События за период_01.10.2023-31.12.2023.xlsx"
        - heat "data/11.Выгрузка_ОДПУ_отопление_ВАО_20240522.xlsx"
        - address "data/13. Адресный реестр объектов недвижимости города Москвы.xlsx"
        Args:
            df (DataFrame): таблица для определения типа

        Returns:
            str: тип таблицы
        """
        file = pd.ExcelFile(file_path)

        # if set(file.sheet_names) == set(HEAT_TABLE_SHEETS):
        #     return "heat"
        if set(file.sheet_names) == set(EVENTS_TABLE_SHEET_NAMES):
            return "events"
        # if set(file.sheet_names) == set(ADDRESS_TABLE_SHEETS):
        #     return "address"

        # if
        return "unknown"

    def __process_event_table(self, s3Key: str, upload_id: int):
        logger.info(f"Downloading table {s3Key}")
        file_path = f"data/{s3Key}"
        if not os.path.exists(file_path):
            self.s3.fget_object("uploads", s3Key, file_path)
        records = process_events_table(file_path)

        table_name = "incident"

        query = """
        insert into incident (description, system, external_created, completed, region_name, unom, address, external_completed, upload_id)
        values (%s, %s, %s, %s, %s, %s, %s, %s, %s);
        """
        chunk_size = 500
        chunks = [
            records[i : i + chunk_size] for i in range(0, len(records), chunk_size)
        ]

        with psycopg.connect(self.dsn) as conn:
            with conn.cursor() as cursor:
                for idx, chunk in enumerate(chunks):

                    values = [record_to_tuple(record, upload_id) for record in chunk]
                    cursor.executemany(query, values)
                    if idx % 10 == 0:
                        logger.info(f"uploaded {idx}/{len(chunks)}")
                    conn.commit()

    def __process_table(self, s3Key: str, upload_id: int):

        logger.info(f"Downloading table {s3Key}")
        file_path = f"data/{s3Key}"
        if not os.path.exists(file_path):
            self.s3.fget_object("uploads", s3Key, file_path)

        table_type = self.__define_table_type(file_path)
        match table_type:
            case "heat":
                logger.info("Processing heat table")
                df = load_heat_df(file_path)
                gen_features_heat_for_group(df)
            case "target":
                logger.info("Processing target table")
                df = load_targets(tuple(file_path))
                process_targets(df)
            case "address":
                logger.info("Processing address table")
                address = load_address(file_path)
                targets = load_processed_targets()
                process_address(address, targets)
            case "events":
                logger.info("Processing events table")
                self.__process_event_table(s3Key, upload_id)
            case "unknown":
                logger.error("Unknown table type")
                return

    def __update_prediction_status(
        self, prediction_id: int, calculated: int, total: int | None = None
    ):
        with psycopg.connect(self.dsn) as conn:
            with conn.cursor() as cur:
                if total is None:
                    query = """
update predictions set calculated = %s, updated_at = %s where id = %s;
"""
                    cur.execute(query, (calculated, dt.datetime.now(), prediction_id))
                else:
                    query = """
update predictions set calculated = %s, updated_at = %s, total = %s where id = %s;
"""
                    cur.execute(
                        query, (calculated, dt.datetime.now(), total, prediction_id)
                    )

    def __predict(self, msg: ConsumerRecord):
        try:
            logger.info("Received prediction message")
            prediction_msg: PredictionRequest = PredictionRequest(
                unoms=msg.value["unoms"],
                start_date=dt.datetime.fromisoformat(msg.value["startDate"]).date(),
                end_date=dt.datetime.fromisoformat(msg.value["endDate"]).date(),
                threshold=msg.value["threshold"],
                region_name=msg.value["regionName"],
                prediction_id=msg.value["predictionID"],
            )

            dates = []
            curDate = prediction_msg["start_date"]
            while curDate <= prediction_msg["end_date"]:
                dates.append(curDate)
                curDate += dt.timedelta(days=1)
            logger.info(
                f'{len(dates)}, {len(prediction_msg["unoms"])}, {prediction_msg["threshold"]}'
            )
            self.__update_prediction_status(
                prediction_msg["prediction_id"],
                0,
                len(prediction_msg["unoms"]) * len(dates),
            )
            chunk_size = 1_000
            unom_splits = [
                prediction_msg["unoms"][i : i + chunk_size]
                for i in range(0, len(prediction_msg["unoms"]), chunk_size)
            ]
            for idx, unom_chunk in enumerate(unom_splits):
                prediction = self.model.predict(
                    unom_chunk, dates, prediction_msg["threshold"]
                )
                self.__save_results(prediction, prediction_msg["region_name"])
                logger.info(
                    f"Prediction: {len(prediction)} in {idx}/{len(unom_splits)}"
                )
                self.__update_prediction_status(
                    prediction_msg["prediction_id"], chunk_size * (idx + 1)
                )
            logger.info("completed predictions")
        except Exception as e:
            logger.error(f"Error: {e}")

    def run(self):
        try:
            for msg in self.consumer:
                headers = self.__parse_headers(msg.headers)
                if "type" in headers and headers["type"] == "predict":
                    logger.info("Received predict message")
                    self.__predict(msg)

                if "type" in headers and headers["type"] == "upload":
                    logger.info("Received upload message")
                    upload_msg: UploadTableRequest = UploadTableRequest(
                        upload_id=msg.value["uploadID"],
                        s3_key=msg.value["s3Key"],
                    )
                    try:
                        self.__update_upload_status(
                            upload_msg["upload_id"], "processing"
                        )
                        self.__process_table(
                            upload_msg["s3_key"], upload_msg["upload_id"]
                        )
                        self.__update_upload_status(upload_msg["upload_id"], "done")
                    except Exception as e:
                        logger.error(f"Error: {e}")
                    finally:
                        self.__update_upload_status(upload_msg["upload_id"], "error")
        except KeyboardInterrupt:
            logger.info("Shutting down")
            self.consumer.close()
            exit(0)


def main():
    dotenv.load_dotenv()
    logger.info("Starting the main function")
    kafka_topic = os.getenv("KAFKA_TOPIC")
    kafka_servers = os.getenv("KAFKA_SERVERS")
    pg_dsn = os.getenv("POSTGRES_DSN")
    minio_host = os.getenv("MINIO_HOST")
    minio_access_key = os.getenv("MINIO_ACCESS_KEY")
    minio_secret_key = os.getenv("MINIO_SECRET_KEY")
    minio_region = os.getenv("MINIO_REGION")
    minio_secure = os.getenv("MINIO_SECURE")

    if not all(
        [
            kafka_topic,
            kafka_servers,
            pg_dsn,
            minio_host,
            minio_access_key,
            minio_secret_key,
            minio_region,
            minio_secure,
        ]
    ):
        raise ValueError("Missing environment variables")

    consumer = KafkaConsumer(
        kafka_topic,
        bootstrap_servers=kafka_servers,
        value_deserializer=lambda m: json.loads(m.decode("utf-8")),
    )

    client = Minio(
        minio_host,  # type: ignore
        access_key=minio_access_key,
        secret_key=minio_secret_key,
        secure=minio_secure.lower() == "true",  # type: ignore
        region=minio_region,
    )

    logger.info("Consumer created")
    model = ModelInference()
    model.load()
    logger.info("Model created")

    controller = Controller(consumer, model, client, pg_dsn)  # type: ignore
    controller.run()


if __name__ == "__main__":
    main()
    logger.info("Ending the main function")
