-- name: CreateTask :one

INSERT INTO tasks (name, created_at, updated_at)
VALUES (?, ?, ?)

RETURNING *;

-- name: GetAllTasks :many

SELECT * FROM tasks;
