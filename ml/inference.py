import pandas as pd


from address import load_processed_address
from heat_features import load_heat_features

import catboost as cb
import datetime as dt

from targets import load_inference_base


class ModelInference:
    def __init__(self):
        self.features = []
        self.model = None

    def load(self):
        self.features.append(load_heat_features().set_index(["unom", "date"]))
        self.features.append(load_processed_address().set_index(["unom", "date"]))
        self.model = cb.CatBoostClassifier()
        self.model.load_model("artifact/model.cbm")

    def predict(self, unoms: list[int], dates: list[dt.date]):
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
        return pd.DataFrame(
            data=out_data, index=base_df[["unom", "date"]], columns=self.model.classes_
        )


if __name__ == "__main__":
    pred = ModelInference()
    pred.load()
    # "..data/dataset/11.Выгрузка_ОДПУ_отопление_ВАО_20240522.xlsx",
    # "..data/dataset/13. Адресный реестр объектов недвижимости города Москвы.xlsx",

    unoms = [302, 16460]
    dates = [dt.date(2024, 4, 6), dt.date(2024, 4, 7)]
    outs = pred.predict(unoms, dates)
    print(outs)
