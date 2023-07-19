// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: permission_roles.sql

package db

import (
	"context"
)

const listAssignedPermissionsToRole = `-- name: ListAssignedPermissionsToRole :many
SELECT permission_id FROM permission_roles
WHERE role_id = ?
`

func (q *Queries) ListAssignedPermissionsToRole(ctx context.Context, roleID int32) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, listAssignedPermissionsToRole, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var permission_id int32
		if err := rows.Scan(&permission_id); err != nil {
			return nil, err
		}
		items = append(items, permission_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}