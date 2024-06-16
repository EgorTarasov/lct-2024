-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists incident(
    id bigserial primary key,
    description text,
    system text,
    external_created timestamp,
    completed timestamp,
    external_completed timestamp,
    region_name text,
    --округ
    unom bigint not null,
    address text not null,
    kind text,
    predicted_at timestamp,
    prediction_id bigint,
    upload_id bigint
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table incident;
-- +goose StatementEnd