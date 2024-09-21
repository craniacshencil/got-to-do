-- name: CreateTask :one
INSERT INTO tasks (task_id, list_id, task_name, start_time, end_time, completion)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetTasks :many
SELECT * from tasks 
WHERE list_id=$1;

-- name: UpdateTask :exec
UPDATE tasks 
SET 
  task_name=$1, 
  start_time=$2, 
  end_time=$3, 
  completion=$4 
WHERE task_id=$5;

-- name: DeleteTask :execrows
DELETE from tasks  
WHERE task_id=$1;
