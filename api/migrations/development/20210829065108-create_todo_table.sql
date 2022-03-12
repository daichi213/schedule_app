
-- +migrate Up
CREATE TABLE todos (
  id bigint NOT NULL,
  created_at timestamp without time zone,
  updated_at timestamp without time zone DEFAULT NULL,
  deleted_at timestamp without time zone DEFAULT NULL,
  title text,
  content text,
  status int DEFAULT NULL,
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS todos;
