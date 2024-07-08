-- name: CreatePassword :one
INSERT INTO user_passwords (id, password)
VALUES ($1, $2)
RETURNING *;

-- name: GetUsernameAndPassword :one
SELECT username, password FROM users NATURAL JOIN user_passwords
WHERE username=$1;
