// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: permission.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createPermission = `-- name: CreatePermission :execresult
INSERT INTO permission (
  permission_name, status, created_at, updated_at
) VALUES (
  ?, ?, ?, ?
)
`

type CreatePermissionParams struct {
	PermissionName string    `json:"permission_name"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (q *Queries) CreatePermission(ctx context.Context, arg CreatePermissionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createPermission,
		arg.PermissionName,
		arg.Status,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const getPermission = `-- name: GetPermission :one
SELECT id, permission_name, status, created_at, updated_at FROM permission
WHERE id = ? LIMIT 1
`

func (q *Queries) GetPermission(ctx context.Context, id int32) (Permission, error) {
	row := q.db.QueryRowContext(ctx, getPermission, id)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.PermissionName,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listPermissions = `-- name: ListPermissions :many
SELECT id, permission_name, status, created_at, updated_at FROM permission
ORDER BY permission_name
`

func (q *Queries) ListPermissions(ctx context.Context) ([]Permission, error) {
	rows, err := q.db.QueryContext(ctx, listPermissions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Permission
	for rows.Next() {
		var i Permission
		if err := rows.Scan(
			&i.ID,
			&i.PermissionName,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePermission = `-- name: UpdatePermission :exec
UPDATE permission
SET permission_name=?, status=?, updated_at=?
WHERE id=?
`

type UpdatePermissionParams struct {
	PermissionName string    `json:"permission_name"`
	Status         string    `json:"status"`
	UpdatedAt      time.Time `json:"updated_at"`
	ID             int32     `json:"id"`
}

func (q *Queries) UpdatePermission(ctx context.Context, arg UpdatePermissionParams) error {
	_, err := q.db.ExecContext(ctx, updatePermission,
		arg.PermissionName,
		arg.Status,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}