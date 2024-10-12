-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS public.problems
(
    id SERIAL NOT NULL,
    title TEXT,
    "description" TEXT,
    specific_location TEXT,
    category TEXT,
    media TEXT[],
    vote_count INT DEFAULT 1,
    lat TEXT NOT NULL,
    long TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (id)
);

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON problems
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_timestamp ON public.problems;

DROP FUNCTION IF EXISTS trigger_set_timestamp();

DROP TABLE IF EXISTS public.problems;
-- +goose StatementEnd
