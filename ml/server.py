import logging
import datetime as dt
import time

import pandas as pd
import grpc
from google.protobuf.timestamp_pb2 import Timestamp
from concurrent import futures

from stubs import inference_pb2_grpc
import stubs.inference_pb2 as pb
from inference import ModelInference
import psycopg

# TODO: загрузка features из s3


def decode_unom_dt(indexStr: str) -> tuple[int, dt.datetime]:
    indexStr = indexStr.replace("(", "").replace(")", "").replace("Timestamp", "")
    unom, date = indexStr.split(",")
    return int(unom), dt.datetime.fromisoformat(date.strip().replace("'", ""))


logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s - %(name)s - %(levelname)s - %(message)s",
)

log = logging.getLogger(__name__)


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


class InferenceService(inference_pb2_grpc.InferenceServicer):
    def __init__(self, pg_dsn: str):
        log.info("Loading model")
        self.dsn = pg_dsn
        self.model = ModelInference()
        self.model.load()
        log.info("Model loaded")

    def __save_results(self, results: dict):
        with psycopg.connect(self.dsn) as conn:
            with conn.cursor() as cur:
                for unom, date, title, value in to_sql_record(results):
                    cur.execute(
                        "insert into incidents(unom, opened_at, predicted_at, title, value, priority) values(%s, %s, %s, %s, %s, %s);",
                        (unom, date, dt.datetime.now(), title, value, 3),
                    )
            conn.commit()

    def Inference(self, request: pb.Query, context):
        unoms = list(request.unoms)

        dates = []
        curTimestamp: Timestamp = request.startDate
        curDate = dt.datetime.fromtimestamp(curTimestamp.seconds)
        endTimestamp: Timestamp = request.endDate
        endDate = dt.datetime.fromtimestamp(endTimestamp.seconds)
        print(endDate, curDate)
        if endDate < curDate:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details("endDate must be greater than startDate")

        while curDate < endDate:
            dates.append(curDate)
            curDate += pd.Timedelta(days=1)
        predictions = []
        try:
            start = time.time()
            outs = self.model.predict(unoms, dates, 0.8)
            self.__save_results(outs)
            log.info(f"predicted in : {time.time() - start}")
        except Exception as e:
            log.error(f"Error: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details("Internal error")
        finally:
            return pb.Response()


if __name__ == "__main__":
    log.info("Starting server")
    dsn = "postgresql://Dino:Dino-misos2024@192.168.1.70:54000/dev"

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    inference_pb2_grpc.add_InferenceServicer_to_server(InferenceService(dsn), server)
    log.info("Server started")
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()
