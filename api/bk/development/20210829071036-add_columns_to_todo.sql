
-- +migrate Up
ALTER TABLE todos
ADD User_id INTEGER NOT NULL;

-- +migrate Down
ALTER TABLE todos
DROP COLUMN User_id;