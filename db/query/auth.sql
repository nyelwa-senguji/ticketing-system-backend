-- name: LoginUser :one
SELECT id, role_id FROM users
WHERE username = ? and password = ?;