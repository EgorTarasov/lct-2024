-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create table if not exists predictions (
    id bigserial primary key,
    user_id bigint,
    adm_area text not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    calculated int default  0,
    total int default 0
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists predictions;
-- +goose StatementEnd
