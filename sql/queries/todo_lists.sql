-- name: CreateList :one
INSERT INTO todo_lists (list_id, user_id, date)
VALUES ($1, $2, $3)
RETURNING *;
