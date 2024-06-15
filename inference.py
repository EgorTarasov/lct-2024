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
        self.features.append(load_heat_features().set_index(['unom', 'date']))
        self.features.append(load_processed_address().set_index(['unom', 'date']))
        self.model = cb.CatBoostClassifier()
        self.model.load_model('artifact/model.cbm')

    def predict(self, unoms: list[int], dates: list[dt.date], no_threshold: float):
        base_df = load_inference_base(unoms, dates)
        base_df = base_df.set_index(['unom', 'date'])
        for feat in self.features:
            base_df = base_df.join(feat, how='left')
        base_df = base_df.reset_index()
        features = base_df.drop(columns=['unom', 'date'])
        pool = cb.Pool(
            features,
            cat_features=features.dtypes[features.dtypes == 'category'].index.tolist()
        )
        out_data = self.model.predict_proba(pool)
        out_data = pd.DataFrame(data=out_data, columns=self.model.classes_)
        out_data[['unom', 'date']] = base_df[['unom', 'date']]
        out_data = out_data[out_data['Нет'] <= no_threshold]
        out_data = out_data.sort_values(by='date').groupby('unom').first().drop(columns='Нет')
        out_data['date'] = out_data['date'].apply(lambda x: x.to_pydatetime().date())
        out_data = out_data.to_dict(orient='index')
        out_data = {k: {'date': v['date'], 'events': {ke: round(ve, 4) for ke, ve in v.items() if ke != 'date' and ve >= 0.01}} for k, v in out_data.items()}
        return out_data


if __name__ == '__main__':
    pred = ModelInference()
    pred.load()
    unoms = [302, 16460]
    dates = [dt.date(2024, 4, x) for x in range(1, 31)]
    outs = pred.predict(unoms, dates, 0.8)
    print(outs)
