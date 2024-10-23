-- +goose Up
-- +goose StatementBegin
ALTER TABLE problems
    ADD COLUMN is_deleted bool not null default false,
    ALTER COLUMN title SET DEFAULT '',
    ALTER COLUMN title SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE problems
    DROP COLUMN is_deleted,
    ADD COLUMN title text;
-- +goose StatementEnd
