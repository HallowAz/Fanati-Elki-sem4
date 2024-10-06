CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS public.users
(
    id SERIAL NOT NULL,
    username TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL,
    birthday TIMESTAMP NOT NULL,
	gender TEXT NOT NULL,
    about TEXT,
    telegram TEXT UNIQUE NOT NULL,
	icon TEXT DEFAULT 'deficon',
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (id),
    CONSTRAINT valid_username CHECK ( LENGTH(USERNAME) >= 3 and LENGTH(USERNAME) <= 20 ),
    CONSTRAINT valid_password CHECK ( LENGTH(PASSWORD) >= 8 and LENGTH(PASSWORD) <= 30 )
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS public.user_rating
(
    user_id INT REFERENCES users(id) NOT NULL,
    count int NOT NULL DEFAULT 0,
	sum_rating int NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (user_id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON user_rating
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS public.user_reviews
(
    id SERIAL NOT NULL,
    "value" smallint NOT NULL,
	reviews TEXT,
    user_id INT REFERENCES users(id) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON user_reviews
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS public.places
(
    id SERIAL NOT NULL,
    title TEXT NOT NULL,
    lat TEXT NOT NULL,
    long TEXT NOT NULL,
    photos TEXT[],
    "description" TEXT,
    group_count INT NOT NULL,
    categories TEXT[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON places
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS public.groups
(
    id SERIAL NOT NULL,
    title TEXT NOT NULL,
	chat_url TEXT,
    people_count INT NOT NULL,
    place_id INT REFERENCES places(id) NOT NULL,
    "date" TIMESTAMP NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON groups
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE IF NOT EXISTS public.user_groups
(
    id SERIAL NOT NULL,
    user_id INT REFERENCES users(id) NOT NULL,
    group_id INT REFERENCES users(id) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (id)
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON user_groups
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
