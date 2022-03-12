
-- +migrate Up
ALTER TABLE todos
ADD user_id INTEGER NOT NULL;

-- +migrate Down
ALTER TABLE todos
DROP COLUMN user_id;