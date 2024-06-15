import os
import dotenv
import logging
from kafka import KafkaConsumer
from typing import TypedDict
import datetime as dt
import json

from inference import ModelInference


dotenv.load_dotenv()


logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)
formatter = logging.Formatter("%(asctime)s - %(name)s - %(levelname)s - %(message)s")
ch = logging.StreamHandler()
ch.setFormatter(formatter)
logger.addHandler(ch)


class PredictionRequest(TypedDict):
    adm_area: str
    start_date: dt.datetime
    end_date: dt.datetime
    threshold: float


def serialize(msg) -> PredictionRequest:
    msg = msg.decode("utf-8")
    data = json.loads(msg)

    return PredictionRequest(
        adm_area=data["admArea"],
        start_date=dt.datetime.fromisoformat(data["startDate"]),
        end_date=dt.datetime.fromisoformat(data["endDate"]),
        threshold=data["threshold"],
    )


class Contoller:

    def __init__(self, consumer: KafkaConsumer, model: ModelInference):
        self.consumer = consumer
        self.model = model

    def run(self):
        for msg in self.consumer:
            try:
                data = msg.value
                adm_area = data["adm_area"]
                start_date = data["start_date"]
                end_date = data["end_date"]
                threshold = data["threshold"]

                prediction = self.model.predict(
                    adm_area, start_date, end_date, threshold
                )
                logger.info(f"Prediction: {prediction}")
                logger.info(f"Prediction type: {type(prediction)}")
                logger.info(f"Prediction shape: {prediction.shape}")
            except Exception as e:
                logger.error(f"Error: {e}")


def main():
    dotenv.load_dotenv()
    logger.info("Starting the main function")
    kafka_topic = os.getenv("KAFKA_TOPIC")
    kafka_servers = os.getenv("KAFKA_SERVERS")
    pg_dsn = os.getenv("PG_DSN")

    consumer = KafkaConsumer(
        "prediction",
        bootstrap_servers="127.0.0.1:9091",
        value_deserializer=serialize,
        auto_offset_reset="earliest",
    )

    logger.info("Consumer created")
    model = ModelInference()
    model.load()
    logger.info("Model created")

    controller = Contoller(consumer, model)
    controller.run()


if __name__ == "__main__":
    main()
    logger.info("Ending the main function")
