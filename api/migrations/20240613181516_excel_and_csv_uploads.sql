-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists uploads(
    id bigserial primary key,
    idempotency_key text, -- uuid
    filename text,
    s3_key text, -- uuid
    status text, -- pending, processing, uploaded
    user_id bigint references users(id),
    created_at timestamp default now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table uploads;
-- +goose StatementEnd
