
-- +migrate Up
CREATE TABLE IF NOT EXISTS schedules (
    id SERIAL PRIMARY KEY,
    created_at timestamp without time zone,
    updated_at timestamp without time zone DEFAULT NULL,
    deleted_at timestamp without time zone DEFAULT NULL,
    title text DEFAULT NULL,
    user_id bigint NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS schedules;