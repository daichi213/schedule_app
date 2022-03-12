
-- +migrate Up
CREATE TABLE IF NOT EXISTS user_schedule_rel (
    id SERIAL PRIMARY KEY,
    created_at timestamp without time zone,
    updated_at timestamp without time zone DEFAULT NULL,
    deleted_at timestamp without time zone DEFAULT NULL,
    user_id bigint NOT NULL,
    schedule_id bigint NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES users (id),
    FOREIGN KEY (schedule_id)
        REFERENCES schedules (id)
);

-- +migrate Down
DROP TABLE IF EXISTS user_schedule_rel;