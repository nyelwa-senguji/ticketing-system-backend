package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
)

type Service interface {
	CreatePermission(ctx context.Context, createPermissionParams db.CreatePermissionParams) (string, error)
	ListPermissions(ctx context.Context) ([]db.Permission, error)
}

type service struct {
	repository *db.Repository
	logger     log.Logger
}

func NewService(repo  *db.Repository, logger log.Logger) Service {
	return &service{
		repository: repo,
		logger:     logger,
	}
}

func (s service) CreatePermission(ctx context.Context, permissionRequest db.CreatePermissionParams) (string, error) {
	logger := log.With(s.logger, "method", "CreatePermission")

	permission := db.CreatePermissionParams{
		PermissionName: permissionRequest.PermissionName,
		Status:         permissionRequest.Status,
	}

	_, err := s.repository.CreatePermission(ctx, permission)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Create Permission", permissionRequest.PermissionName)

	return "Permission Created Succesfully", nil
}

func (s service) ListPermissions(ctx context.Context) ([]db.Permission, error) {
	logger := log.With(s.logger, "method", "ListPermission")

	var permissions []db.Permission

	permissions, err := s.repository.ListPermissions(ctx)

	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	return permissions, nil
}
