-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.users
(
    id SERIAL NOT NULL,
    username TEXT,
    phone TEXT,
    icon TEXT,
    password TEXT,
    age INT,
    gender TEXT,
    is_admin BOOLEAN,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (id)
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON users
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_timestamp ON public.users;

DROP FUNCTION IF EXISTS trigger_set_timestamp();

DROP TABLE IF EXISTS public.users;
-- +goose StatementEnd
