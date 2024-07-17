-- name: GetTodo :one
SELECT * FROM todos
WHERE id = ? LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY name;

-- name: CreateTodo :exec
INSERT INTO todos (
  name
) VALUES (
  ?
);

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;