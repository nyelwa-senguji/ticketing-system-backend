package service

import (
	"context"
	"reflect"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
)

type Service interface {
	CreatePermission(ctx context.Context, createPermissionParams CreatePermissionRequest) (string, error)
	ListPermissions(ctx context.Context) ([]db.Permission, error)
	GetPermission(ctx context.Context, id int32) (db.Permission, error)
	UpdatePermission(ctx context.Context, updateReq UpdatePermissionRequest) (string, error)
}

type CreatePermissionRequest struct {
	PermissionName string `json:"permission_name"`
	Status         string `json:"status"`
}

type UpdatePermissionRequest struct{
	PermissionName string    `json:"permission_name"`
	Status         string    `json:"status"`
	ID             int32     `json:"id"`
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

func (s service) CreatePermission(ctx context.Context, permissionRequest CreatePermissionRequest) (string, error) {

	logger := log.With(s.logger, "method", "CreatePermission")

	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

	permission := db.CreatePermissionParams{
		PermissionName: permissionRequest.PermissionName,
		Status:         permissionRequest.Status,
		CreatedAt:      time,
		UpdatedAt:      time,
	}

	if reflect.DeepEqual(permission.PermissionName, "") {
		return "Permission name can not be empty", nil
	}

	if reflect.DeepEqual(permission.Status, "") {
		return "Status can not be empty", nil
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

func (s service) GetPermission(ctx context.Context, id int32) (db.Permission, error) {
	logger := log.With(s.logger, "method", "GetPermission")

	var permission db.Permission

	permission, err := s.repository.GetPermission(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return permission, err
	}

	return permission, nil
}

func (s service) UpdatePermission(ctx context.Context, updateReq UpdatePermissionRequest) (string, error) {
	logger := log.With(s.logger, "method", "UpdatePermission")

	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

	updatePermission := db.UpdatePermissionParams{
		PermissionName: updateReq.PermissionName,
		Status: updateReq.Status,
		UpdatedAt: time,
		ID: updateReq.ID,
	}

	err := s.repository.UpdatePermission(ctx, updatePermission)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "",err
	}

	return  "Permission updated successfully",nil
}
