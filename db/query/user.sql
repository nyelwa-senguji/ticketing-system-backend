-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: LoginUser :one
SELECT id, role_id FROM users
WHERE username = ? and password = ?;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :execresult
INSERT INTO users (
  username, password, created_at, updated_at, role_id
) VALUES (
  ?, ?, ?, ?, ?
);
