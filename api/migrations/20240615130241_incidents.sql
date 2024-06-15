-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists incidents (
    id bigserial primary key,
    opened_at timestamp,
    predicted_at timestamp,
    closed_at timestamp default  null,
    title text not null,
    priority int not null,
    status text not null default 'active',
    unom bigint,
    created_at timestamp default now(),
    updated_at timestamp default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table incidents;
-- +goose StatementEnd
