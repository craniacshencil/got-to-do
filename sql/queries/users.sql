-- name: CreateUser :one
INSERT INTO
  users (id, username, first_name, last_name)
VALUES( $1, $2, $3, $4 )
RETURNING *;

-- name: GetUsername :one
SELECT * FROM users WHERE username=$1;
