// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package db

import (
	"time"
)

type Category struct {
	ID           int32     `json:"id"`
	CategoryName string    `json:"category_name"`
	Status       string    `json:"status"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Permission struct {
	ID             int32     `json:"id"`
	PermissionName string    `json:"permission_name"`
	Status         string    `json:"status"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}

type PermissionRoles struct {
	PermissionID int32 `json:"permission_id"`
	RoleID       int32 `json:"role_id"`
}

type Roles struct {
	ID        int32     `json:"id"`
	RoleName  string    `json:"role_name"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Tickets struct {
	ID          int32     `json:"id"`
	Subject     string    `json:"subject"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      int32     `json:"user_id"`
	CategoryID  int32     `json:"category_id"`
}

type Users struct {
	ID        int32     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	RoleID    int32     `json:"role_id"`
}
