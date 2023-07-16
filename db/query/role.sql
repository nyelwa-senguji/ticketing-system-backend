-- name: GetRole :one
SELECT * FROM roles
WHERE id = ? LIMIT 1;

-- name: GetRoleByName :one
SELECT * FROM roles
WHERE role_name = ? LIMIT 1;

-- name: ListRoles :many
SELECT * FROM roles
ORDER BY role_name;

-- name: CreateRole :execresult
INSERT INTO roles (
  role_name, status, created_at, updated_at
) VALUES (
  ?, ?, ?, ?
);

-- name: UpdateRole :exec
UPDATE roles
SET role_name=?, status=?, updated_at=?
WHERE id=?