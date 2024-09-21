-- name: CreateList :one
INSERT INTO todo_lists (list_id, user_id, date)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetListID :one
SELECT list_id from todo_lists
WHERE date=$1 and user_id=$2;

-- name: DeleteList :execrows
DELETE FROM todo_lists
WHERE list_id=$1;

