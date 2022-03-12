
-- +migrate Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at timestamp without time zone,
    updated_at timestamp without time zone DEFAULT NULL,
    deleted_at timestamp without time zone DEFAULT NULL,
    email text,
    pass bytea
);

-- +migrate Down
DROP TABLE IF EXISTS users;
