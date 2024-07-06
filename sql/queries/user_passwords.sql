-- name: CreatePassword :one
INSERT INTO user_passwords (id, password)
VALUES ($1, $2)
RETURNING *;
