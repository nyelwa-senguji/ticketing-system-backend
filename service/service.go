package service

import (
	"context"
	"github.com/go-kit/kit/log"

	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
)

type Service interface {
	/******************************
		Services for Permission
	*******************************/
	CreatePermission(ctx context.Context, createPermissionParams CreatePermissionRequest) (string, error)
	ListPermissions(ctx context.Context) ([]db.Permission, error)
	GetPermission(ctx context.Context, id int32) (db.Permission, error)
	UpdatePermission(ctx context.Context, updateReq UpdatePermissionRequest) (string, error)

	/******************************
		Services for Roles
	*******************************/
	CreateRole(ctx context.Context, createRoleReq CreateRoleRequest) (string, error)
	ListRoles(ctx context.Context) ([]db.Roles, error)
	GetRole(ctx context.Context, id int32) (db.Roles, error)
	UpdateRole(ctx context.Context, updateRoleReq UpdateRoleRequest)(string, error)
}

type service struct {
	repository *db.Repository
	logger     log.Logger
}

func NewService(repo *db.Repository, logger log.Logger) Service {
	return &service{
		repository: repo,
		logger:     logger,
	}
}