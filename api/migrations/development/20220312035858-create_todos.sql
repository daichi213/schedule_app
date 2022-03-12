
-- +migrate Up
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    created_at timestamp without time zone,
    updated_at timestamp without time zone DEFAULT NULL,
    deleted_at timestamp without time zone DEFAULT NULL,
    schedule_id bigint NOT NULL,
    FOREIGN KEY (schedule_id)
        REFERENCES schedules (id)
);

-- +migrate Down
DROP TABLE IF EXISTS todos;
