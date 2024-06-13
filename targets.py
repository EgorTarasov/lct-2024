import os

import pandas as pd
import datetime as dt

GOOD_CLASSES = {
    'P1 <= 0',
    'P2 <= 0',
    'T1 > max',
    'T1 < min',
    'Авария',
    'Недостаточная температура подачи в центральном отоплении (Недотоп)',
    'Превышение температуры подачи в центральном отоплении (Перетоп)',
    'Утечка теплоносителя',
    'Течь в системе отопления',
    'Температура в квартире ниже нормативной',
    'Отсутствие отопления в доме',
    'Сильная течь в системе отопления',
    'Крупные пожары',
    'Температура в помещении общего пользования ниже нормативной',
    'Протечка труб в подъезде'
}


def load_targets():
    df1 = pd.read_excel('data/События за период_01.01.2024-30.04.2024.xlsx', sheet_name='Выгрузка')
    df2 = pd.read_excel('data/События за период_01.10.2023-31.12.2023.xlsx', sheet_name='Выгрузка')
    df = pd.concat((df1, df2)).reset_index(drop=True)
    df = df[~df['УНОМ'].isna()]
    df = df[df['Наименование'].isin(GOOD_CLASSES)]
    df['Дата создания во внешней системе'] = pd.to_datetime(df['Дата создания во внешней системе'])
    df['Дата закрытия'] = pd.to_datetime(df['Дата закрытия'])
    df['Дата и время завершения события во внешней системе'] = pd.to_datetime(
        df['Дата и время завершения события во внешней системе'])
    df['date'] = df['Дата создания во внешней системе'].apply(lambda x: x.date())
    df['УНОМ'] = df['УНОМ'].astype(int)
    df.to_parquet('features/target_pre.parquet')
    return df


def process_targets(df: pd.DataFrame):
    all_unoms = df['УНОМ'].unique()
    all_dates = pd.date_range(df['date'].min(), df['date'].max())

    df_t = df.groupby(['УНОМ', 'date'])['Наименование'].unique()

    target_order = df_t.explode().dropna().value_counts().index.tolist()

    def find_first_target(targets: list[str]) -> str:
        indices = [target_order.index(x) for x in targets]
        return target_order[min(indices)]

    df_t = df_t.apply(find_first_target)

    all_index = pd.MultiIndex.from_product([all_unoms, all_dates], names=['unom', 'date'])
    df_t = df_t.reindex(all_index, fill_value='Нет')
    df_t = df_t.reset_index()
    df_t['Наименование'] = df_t['Наименование'].astype('category')
    df_t['month'] = df_t['date'].apply(lambda x: x.month).astype('category')
    df_t['day_of_week'] = df_t['date'].apply(lambda x: x.dayofweek).astype('category')
    return df_t


def load_processed_targets():
    df_t = pd.read_parquet('features/target.parquet')
    df_t['month'] = df_t['month'].astype('category')
    df_t['day_of_week'] = df_t['day_of_week'].astype('category')
    return df_t


def load_inference_base(unoms: list[int], dates: list[dt.date]):
    all_index = pd.MultiIndex.from_product((unoms, dates), names=['unom', 'date'])
    df = pd.DataFrame(index=all_index)
    df = df.reset_index()
    df['date'] = pd.to_datetime(df['date'])
    df['month'] = df['date'].apply(lambda x: x.month).astype('category')
    df['day_of_week'] = df['date'].apply(lambda x: x.dayofweek).astype('category')
    return df


if __name__ == '__main__':
    tdf = load_targets()
    fts = process_targets(tdf)
