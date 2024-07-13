-- name: CreateTask :one
INSERT INTO tasks (task_id, list_id, task_name, start_time, end_time)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
