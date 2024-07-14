-- name: CreateTask :one
INSERT INTO tasks (task_id, list_id, task_name, start_time, end_time, completion)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetTasks :many
SELECT * from tasks 
WHERE list_id=$1;
