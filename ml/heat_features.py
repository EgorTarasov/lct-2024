import os.path

import pandas as pd
from tqdm.contrib.concurrent import process_map


def load_heat_df(
    df_path: str = "data/11.Выгрузка_ОДПУ_отопление_ВАО_20240522.xlsx",
):
    df_heat_1 = pd.read_excel(df_path, sheet_name="Sheet 1")
    df_heat_2 = pd.read_excel(df_path, sheet_name="Sheet 2")
    df_heat = pd.concat((df_heat_1, df_heat_2))
    df_heat["UNOM"] = df_heat["UNOM"].astype(int)
    df_heat["date"] = pd.to_datetime(df_heat["Месяц/Год"], format="%d-%m-%Y")
    df_heat["Объём поданого теплоносителя в систему ЦО"] = (
        df_heat["Объём поданого теплоносителя в систему ЦО"].fillna(0).astype(float)
    )
    df_heat["Объём обратного теплоносителя из системы ЦО"] = (
        df_heat["Объём обратного теплоносителя из системы ЦО"].fillna(0).astype(float)
    )
    df_heat["Объём обратного теплоносителя из системы ЦО"] = (
        df_heat["Объём обратного теплоносителя из системы ЦО"].fillna(0).astype(float)
    )
    df_heat["Разница между подачей и обраткой(Подмес)"] = (
        df_heat["Разница между подачей и обраткой(Подмес)"].fillna(0).astype(float)
    )
    df_heat["Разница между подачей и обраткой(Утечка)"] = (
        df_heat["Разница между подачей и обраткой(Утечка)"].fillna(0).astype(float)
    )
    df_heat["Температура подачи"] = df_heat["Температура подачи"].astype(float)
    df_heat["Температура обратки"] = df_heat["Температура обратки"].astype(float)
    df_heat["Разница температур"] = (
        df_heat["Температура подачи"] - df_heat["Температура обратки"]
    )
    df_heat["Наработка часов счётчика"] = (
        df_heat["Наработка часов счётчика"].fillna(0).astype(float)
    )
    df_heat["Расход тепловой энергии "] = (
        df_heat["Расход тепловой энергии "].fillna(0).astype(float)
    )
    df_heat["Ошибки"] = df_heat["Ошибки"].str.split(",").fillna("").apply(set)
    df_heat.to_parquet("features/heat_pre.parquet")
    return df_heat


def gen_features_heat_for_group(group):
    rolls = [3, 7, 30]
    lags = [1, 2, 3]
    agg_features = [
        "Объём поданого теплоносителя в систему ЦО",
        "Объём обратного теплоносителя из системы ЦО",
        "Разница между подачей и обраткой(Подмес)",
        "Разница между подачей и обраткой(Утечка)",
        "Температура подачи",
        "Температура обратки",
        "Разница температур",
        "Наработка часов счётчика",
        "Расход тепловой энергии ",
    ]
    agg_funs = [
        "sum",
        "mean",
        "max",
        "min",
        percentile(0.9),
        percentile(0.1),
        percentile(0.5),
        percentile(0.75),
        percentile(0.25),
    ]
    to_cat = []
    group = group.sort_values(by="date")
    for roll in rolls:
        for feature in agg_features:
            group_roll = group.rolling(roll)[feature].agg(agg_funs).shift(1)
            group_roll.columns = [
                f"roll_{roll}_feature_{feature}_{x}" for x in group_roll.columns
            ]
            to_cat.append(group_roll)
    for lag in lags:
        group_lag = group[agg_features].shift(lag)
        group_lag.columns = [f"lag_{lag}_feature_{x}" for x in group_lag.columns]
        to_cat.append(group_lag)
    catted = pd.concat(to_cat, axis=1)
    catted["date"] = group["date"]
    catted["unom"] = group["UNOM"]
    return catted


def percentile(n):
    def percentile_(x):
        return x.quantile(n)

    percentile_.__name__ = "percentile_{:02.0f}".format(n * 100)
    return percentile_


def gen_heat_features(df_heat):
    all_unoms = df_heat["UNOM"].unique()
    results = process_map(
        gen_features_heat_for_group,
        [df_heat[df_heat["UNOM"] == unom].copy() for unom in all_unoms],
        chunksize=2,
        max_workers=30,
    )
    df_heat_features = pd.concat(results)
    df_heat_features.to_parquet("features/heat.parquet")


def load_heat_features():
    return pd.read_parquet("features/heat.parquet")


if __name__ == "__main__":
    df = load_heat_df("../data/dataset/11.Выгрузка_ОДПУ_отопление_ВАО_20240522.xlsx")
    gen_heat_features(df)
