import pandas as pd
from sklearn.model_selection import train_test_split

from address import load_processed_address
from heat_features import load_heat_features
from targets import load_processed_targets
import catboost as cb


def make_df():
    targets = load_processed_targets()
    address = load_processed_address()
    heat = load_heat_features()
    df_t = targets.set_index(['unom', 'date']).join(heat.set_index(['unom', 'date']), how='left').join(
        address.set_index(['unom', 'date']), how='left')
    df_t = df_t.reset_index()
    return df_t


def make_pool(data):
    features = data.drop(columns=['unom', 'date', 'Наименование'])
    return cb.Pool(
        features,
        data['Наименование'],
        cat_features=features.dtypes[features.dtypes == 'category'].index.tolist()
    )


if __name__ == '__main__':
    target_df = make_df()
    df_t_train, df_t_test = train_test_split(target_df, test_size=0.1, stratify=target_df['Наименование'])
    pool_train, pool_test = make_pool(df_t_train), make_pool(df_t_test)
    cbs = cb.CatBoostClassifier(
        loss_function='MultiClass',
        auto_class_weights='SqrtBalanced',
        eval_metric='TotalF1:average=Macro;use_weights=false',
        task_type='GPU',
        iterations=1000,
        random_state=0xDEADBEEF
    )
    cbs.fit(pool_train, eval_set=pool_test)
    cbs.save_model('artifact/model.cbm')
    cbs.get_feature_importance(pool_test, prettified=True).to_parquet('artifact/feature_importance.parquet')
