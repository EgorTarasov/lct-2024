import datetime as dt
from typing import NamedTuple
import logging

import time

import psycopg
import pandas as pd
from minio import Minio


from heat_features import load_heat_df, gen_features_heat_for_group
from address import load_address, process_address
from targets import load_targets, process_targets, load_processed_targets


logging.basicConfig(
    level=logging.INFO, format="%(asctime)s - %(name)s - %(levelname)s - %(message)s"
)

log = logging.getLogger(__name__)


class Upload(NamedTuple):
    id: int
    s3_key: str
    created_at: dt.datetime


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


class ProcessorService:
    def __init__(self, dsn: str, s3_client: Minio):
        self.dsn = dsn
        self.s3 = s3_client
        self.conn = None
        self.queue = []

    def connect(self):
        self.conn = psycopg.connect(self.dsn)

    def disconnect(self):
        if self.conn is not None:
            self.conn.close()
        self.conn = None

    def __fetch_unprocessed(self) -> list[Upload]:
        with psycopg.connect(self.dsn) as conn:

            with conn.cursor() as cur:
                print("fetching", cur)
                cur.execute(
                    """select 
                                u.id,
                                u.s3_key,
                                u.created_at
                            from uploads u
                            where status = 'pending'
                            order by u.created_at;
                        """
                )
                return [Upload(x[0], x[1], x[2]) for x in cur.fetchall()]

    def __update_table_status(self, upload_id: int, status: str):
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

    def __define_table_type(self, file_path) -> str:
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

        if set(file.sheet_names) == set(HEAT_TABLE_SHEETS):
            return "heat"
        if set(file.sheet_names) == set(EVENTS_TABLE_SHEET_NAMES):
            return "target"
        if set(file.sheet_names) == set(ADDRESS_TABLE_SHEETS):
            return "address"
        return "unknown"

    def __process_table(self, s3Key: str):

        log.info(f"Downloading table {s3Key}")
        file_path = f"data/{s3Key}"
        self.s3.fget_object("uploads", s3Key, file_path)

        table_type = self.__define_table_type(file_path)
        match table_type:
            case "heat":
                log.info("Processing heat table")
                df = load_heat_df(file_path)
                gen_features_heat_for_group(df)
            case "target":
                log.info("Processing target table")
                df = load_targets(tuple(file_path))
                process_targets(df)
            case "address":
                log.info("Processing address table")
                address = load_address(file_path)
                targets = load_processed_targets()
                process_address(address, targets)
            case "unknown":
                log.error("Unknown table type")
                return

    def run(self):

        while True:
            log.info("Checking for new uploads")

            try:
                new_uploads = self.__fetch_unprocessed()
                log.info(f"got new uploads {len(new_uploads)}")
                if len(new_uploads) == 0 and len(self.queue) == 0:
                    time.sleep(60)
                    continue

                for upload in new_uploads:
                    if upload.s3_key in self.queue:
                        continue

                    self.queue.append(upload)
            except Exception as e:
                log.error(f"Error: {e}")

            start = time.time()
            upload_id = None
            try:
                upload = self.queue.pop(0)
                upload_id = upload.id
                self.__update_table_status(upload.id, "processing")
                self.__process_table(upload.s3_key)
                log.info(f"Processed {upload.id} in {time.time() - start} seconds")
                self.__update_table_status(upload.id, "done")
            except Exception as e:
                log.error(f"Error: {e}")
                if upload_id is not None:
                    self.__update_table_status(upload.id, "error")


if __name__ == "__main__":

    client = Minio(
        "localhost:9000",
        access_key="key",
        secret_key="key",
        secure=False,
        region="eu-west-1",
    )
    dsn = "postgresql://user:passord@localhost:5432/dev"
    log.info(f"pg connection {dsn}")

    service = ProcessorService(dsn, client)
    log.info("Starting service")
    service.run()
