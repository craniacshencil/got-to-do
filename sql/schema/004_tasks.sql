-- +goose Up
CREATE TABLE tasks (
  task_id UUID PRIMARY KEY,
  list_id UUID NOT NULL REFERENCES todo_lists,
  task_name VARCHAR(150) NOT NULL,
  start_time TIME
  WITH
    TIME ZONE NOT NULL,
    end_time TIME
  WITH
    TIME ZONE NOT NULL,
    completion BOOLEAN NOT NULL DEFAULT FALSE
);

-- +goose Down
DROP TABLE tasks;
