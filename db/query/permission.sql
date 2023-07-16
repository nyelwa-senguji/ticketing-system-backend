-- name: GetPermission :one
SELECT * FROM permission
WHERE id = ? LIMIT 1;

-- name: GetPermissionByName :one
SELECT * FROM permission
WHERE permission_name = ? LIMIT 1;

-- name: ListPermissions :many
SELECT * FROM permission
ORDER BY permission_name;

-- name: CreatePermission :execresult
INSERT INTO permission (
  permission_name, status, created_at, updated_at
) VALUES (
  ?, ?, ?, ?
);

-- name: UpdatePermission :exec
UPDATE permission
SET permission_name=?, status=?, updated_at=?
WHERE id=?