-- name: ListAssignedPermissionsToRole :many
SELECT permission_id FROM permission_roles
WHERE role_id = ?;

-- name: AssignPermissionToRole :one
INSERT INTO permission_roles (permission_id, role_id)
VALUES(?, ?)