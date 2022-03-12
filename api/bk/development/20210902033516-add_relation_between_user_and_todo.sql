
-- +migrate Up
ALTER TABLE todos
ADD CONSTRAINT todo_ibfk_1 FOREIGN KEY (User_id) REFERENCES users (id);

-- +migrate Down
ALTER TABLE todos
DROP FOREIGN KEY todo_ibfk_1;
