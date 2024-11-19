-- +goose Up
-- +goose StatementBegin
ALTER TABLE problems
    ADD COLUMN status TEXT NOT NULL DEFAULT 'Ожидание';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE problems
    DROP COLUMN status;
-- +goose StatementEnd
