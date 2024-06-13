import pandas as pd

from targets import load_processed_targets


def load_address():
    df_address = pd.read_excel('data/13. Адресный реестр объектов недвижимости города Москвы.xlsx', sheet_name='0')
    df_address = df_address.iloc[1:]
    df_address['UNOM'] = df_address['UNOM'].astype(int)
    df_address = df_address.set_index('UNOM')
    df_address = df_address[~df_address.index.duplicated(keep='first')]
    df_address['DREG'] = pd.to_datetime(df_address['DREG'], format='%d.%m.%Y')
    df_address['DISTRICT'] = df_address['DISTRICT'].astype('category')
    df_address['ADM_AREA'] = df_address['ADM_AREA'].astype('category')
    df_address.to_parquet('features/address_pre.parquet')
    return df_address


def process_address(df_address, df_target):
    df_target = df_target.set_index('unom').join(df_address[['ADM_AREA', 'DISTRICT', 'DREG']], how='left').reset_index()
    df_target['reg_years_ago'] = (df_target['date'] - df_target['DREG']).apply(
        lambda x: x.total_seconds() / 60 / 60 / 24 / 30 / 12
    ).fillna(-1)

    df = df_target[['unom', 'date', 'reg_years_ago']]
    df.to_parquet('features/address.parquet')


def load_processed_address():
    return pd.read_parquet('features/address.parquet')


if __name__ == '__main__':
    address = load_address()
    targets = load_processed_targets()
    process_address(address, targets)
