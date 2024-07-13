-- name: CreatePassword :one
INSERT INTO user_passwords (id, password)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserDetails :one
SELECT users.id, username, password FROM users NATURAL JOIN user_passwords
WHERE username=$1;
