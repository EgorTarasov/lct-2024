-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE incidents
ADD COLUMN value FLOAT;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE incidents DROP COLUMN value;
-- +goose StatementEnd