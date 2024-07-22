-- name: GetTodo :one
SELECT * FROM todos
WHERE id = ? LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
WHERE user_id = ?
ORDER BY name;

-- name: CreateTodo :one
INSERT INTO todos (
  user_id, name
) VALUES (
  ?, ?
)
RETURNING id;
;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;