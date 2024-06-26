import pandas as pd

from address import load_processed_address
from heat_features import load_heat_features

import catboost as cb
import datetime as dt

from typing import TypedDict


from targets import load_inference_base


# class PredictRecord(TypedDict):
# Define a dictionary with keys containing Russian letters and special characters
my_dict = {
    "P1 <= 0": True,
    "P2 <= 0": False,
    "T1 < мин": True,  # Russian word for "min"
    "T1 > макс": False,  # Russian word for "max"
    "Нет": True,
    "Отсутствие отопления в доме": False,
    "Протечка труб в подъезде": True,
    "Сильная течь в системе отопления": False,
    "Температура в квартире ниже нормативной": True,
    "Температура в помещении общего пользования ниже нормативной": False,
    "Течь в системе отопления": True,
}

# List the dictionary, printing each key-value pair
for key, value in my_dict.items():
    print(f"{key}: {value}")

class ModelInference:
    def __init__(self):
        self.features = []
        self.model = None

    def load(self):
        self.features.append(load_heat_features().set_index(["unom", "date"]))
        self.features.append(load_processed_address().set_index(["unom", "date"]))
        self.model = cb.CatBoostClassifier()
        self.model.load_model("artifact/model.cbm")

    def predict(
        self, unoms: list[int], dates: list[dt.date], no_threshold: float
    ) -> list[dict]:
        base_df = load_inference_base(unoms, dates)
        base_df = base_df.set_index(["unom", "date"])
        for feat in self.features:
            base_df = base_df.join(feat, how="left")
        base_df = base_df.reset_index()
        features = base_df.drop(columns=["unom", "date"])
        pool = cb.Pool(
            features,
            cat_features=features.dtypes[features.dtypes == "category"].index.tolist(),
        )
        out_data = self.model.predict_proba(pool)
        out_data = pd.DataFrame(data=out_data, columns=self.model.classes_)
        out_data = out_data[out_data["Нет"] <= no_threshold]
        # remove index from out_data
        out_data = out_data.reset_index(drop=True)
        out_data[["unom", "date"]] = base_df[["unom", "date"]]
        records = out_data.to_dict(orient="records")
        return records
        # out_data[["unom", "date"]] = base_df[["unom", "date"]]
        # out_data = out_data[out_data["Нет"] <= no_threshold]
        # out_data = (
        #     out_data.sort_values(by="date").groupby("unom").first().drop(columns="Нет")
        # )
        # out_data["date"] = out_data["date"].apply(lambda x: x.to_pydatetime().date())
        # out_data = out_data.to_dict(orient="index")
        # out_data = {
        #     k: {
        #         "date": v["date"],
        #         "events": {
        #             ke: round(ve, 4)
        #             for ke, ve in v.items()
        #             if ke != "date" and ve >= 0.01
        #         },
        #     }
        #     for k, v in out_data.items()
        # }
        # return out_data


if __name__ == "__main__":
    pred = ModelInference()
    pred.load()
    # "..data/dataset/11.Выгрузка_ОДПУ_отопление_ВАО_20240522.xlsx",
    # "..data/dataset/13. Адресный реестр объектов недвижимости города Москвы.xlsx",
    unoms = [302, 16460]
    dates = [dt.date(2024, 4, x) for x in range(1, 31)]
    outs = pred.predict(unoms, dates, 0.8)
    print(outs)
