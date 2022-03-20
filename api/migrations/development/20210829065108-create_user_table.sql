
-- +migrate Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at timestamp without time zone,
    updated_at timestamp without time zone DEFAULT NULL,
    deleted_at timestamp without time zone DEFAULT NULL,
    username text NOT NULL,
    email text NOT NULL,
    password bytea NOT NULL,
    admin_flag int NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS users;
