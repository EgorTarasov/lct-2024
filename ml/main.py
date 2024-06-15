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

# TODO: загрузка features из s3


def decode_unom_dt(indexStr: str) -> tuple[int, dt.datetime]:
    indexStr = indexStr.replace("(", "").replace(")", "").replace("Timestamp", "")
    unom, date = indexStr.split(",")
    return int(unom), dt.datetime.fromisoformat(date.strip().replace("'", ""))


def record_to_pb(record: tuple) -> pb.Prediction:
    unom, date = record[0]
    return pb.Prediction(
        unom=unom,
        date=Timestamp(seconds=int(date.timestamp())),
        p1=record[1],
        p2=record[2],
        t1=record[3],
        t2=record[4],
        no=record[5],
        noHeating=record[6],
        leak=record[7],
        strongLeak=record[8],
        tempLow=record[9],
        tempLowCommon=record[10],
        leakSystem=record[11],
    )


logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s - %(name)s - %(levelname)s - %(message)s",
)

log = logging.getLogger(__name__)


class InferenceService(inference_pb2_grpc.InferenceServicer):
    def __init__(self):
        log.info("Loading model")
        self.model = ModelInference()
        self.model.load()
        log.info("Model loaded")

    def Inference(self, request: pb.Query, context):
        # TODO: try except
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
            return pb.Response(predictions=[])

        while curDate < endDate:
            dates.append(curDate)
            curDate += pd.Timedelta(days=1)
        predictions = []
        try:
            start = time.time()
            outs = self.model.predict(unoms, dates)
            df_tuples = [tuple(row) for row in outs.to_records(index=True)]
            predictions = [record_to_pb(record) for record in df_tuples]
            log.info(f"predicted in : {time.time() - start}")
        except Exception as e:
            log.error(f"Error: {e}")
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details("Internal error")
        finally:
            return pb.Response(predictions=predictions)


if __name__ == "__main__":
    log.info("Starting server")
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    inference_pb2_grpc.add_InferenceServicer_to_server(InferenceService(), server)
    log.info("Server started")
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()
