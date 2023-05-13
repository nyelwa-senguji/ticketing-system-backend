-- name: GetRole :one
SELECT * FROM roles
WHERE id = ? LIMIT 1;

-- name: ListRoles :many
SELECT * FROM roles
ORDER BY role_name;

-- name: CreateRole :execresult
INSERT INTO roles (
  role_name, status, created_at, updated_at
) VALUES (
  ?, ?, ?, ?
);

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = ?;