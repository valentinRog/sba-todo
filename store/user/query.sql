-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users
WHERE username = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  username, password
) VALUES (
  ?, ?
) RETURNING id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;