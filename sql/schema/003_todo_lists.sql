-- +goose Up
CREATE TABLE todo_lists (
  list_id UUID PRIMARY KEY,
  user_id UUID REFERENCES users (id),
  date DATE NOT NULL
);

-- +goose Down
DROP TABLE todo_lists;
