-- name: ListTodos :many
SELECT * FROM todos 
LIMIT 10;

-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 
LIMIT 1;

-- name: CreateTodo :one
INSERT INTO todos (todo, description)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
SET todo = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos 
WHERE id = $1;
