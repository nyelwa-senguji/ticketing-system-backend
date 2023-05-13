-- name: GetPermission :one
SELECT * FROM permission
WHERE id = ? LIMIT 1;

-- name: ListPermissions :many
SELECT * FROM permission
ORDER BY permission_name;

-- name: CreatePermission :execresult
INSERT INTO permission (
  permission_name, status, created_at, updated_at
) VALUES (
  ?, ?, ?, ?
);

-- name: DeletePermission :exec
DELETE FROM permission
WHERE id = ?;