-- +goose Up
CREATE TABLE tasks (
  task_id UUID PRIMARY KEY,
  list_id UUID NOT NULL REFERENCES todo_lists,
  task_name VARCHAR(150) NOT NULL,
  start_time TIME NOT NULL,
  end_time TIME NOT NULL
);

-- +goose Down
DROP TABLE tasks;
