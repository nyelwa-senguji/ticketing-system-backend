package service

import (
	"context"
	"reflect"

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
	/*********************************************************
		1. Logging the method name
	*********************************************************/
	logger := log.With(s.logger, "method", "CreatePermission")

	/*********************************************************
		2. Assigning permision variable with new request
	*********************************************************/
	permission := db.CreatePermissionParams{
		PermissionName: permissionRequest.PermissionName,
		Status:         permissionRequest.Status,
		CreatedAt: permissionRequest.CreatedAt,
		UpdatedAt: permissionRequest.UpdatedAt,
	}

	/*********************************************************
		3. Check if permission name is empty
	*********************************************************/
	if(reflect.DeepEqual(permission.PermissionName, "")){
		return "Permission name can not be empty", nil
	}

	/*********************************************************
		4. Check if status is empty
	*********************************************************/
	if(reflect.DeepEqual(permission.Status, "")){
		return "Status can not be empty", nil
	}

	/*********************************************************
		5. if Ok, submitting permission to the database
	*********************************************************/
	_, err := s.repository.CreatePermission(ctx, permission)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	/*********************************************************
		6. Logging the results
	*********************************************************/
	logger.Log("Create Permission", permissionRequest.PermissionName)

	/*********************************************************
		7. Returning response
	*********************************************************/
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
