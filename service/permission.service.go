package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/log/level"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
)

type CreatePermissionRequest struct {
	PermissionName string `json:"permission_name"`
	Status         string `json:"status"`
}

type UpdatePermissionRequest struct {
	PermissionName string `json:"permission_name"`
	Status         string `json:"status"`
	ID             int32  `json:"id"`
}

func (s service) CreatePermission(ctx context.Context, permissionRequest CreatePermissionRequest) (string, error) {

	logger := log.With(s.logger, "method", "CreatePermission")

	checkIfPermissionExists, _ := s.repository.GetPermissionByName(ctx, permissionRequest.PermissionName)

	if(checkIfPermissionExists.PermissionName == permissionRequest.PermissionName){
		return "Permission already exists", nil
	}

	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

	permission := db.CreatePermissionParams{
		PermissionName: permissionRequest.PermissionName,
		Status:         permissionRequest.Status,
		CreatedAt:      time,
		UpdatedAt:      time,
	}

	_, err := s.repository.CreatePermission(ctx, permission)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Create Permission", permission.PermissionName)

	return "Permission Created Succesfully", nil
}

func (s service) ListPermissions(ctx context.Context) ([]db.Permission, error) {
	logger := log.With(s.logger, "method", "ListPermission")

	permissions, err := s.repository.ListPermissions(ctx)

	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	logger.Log("List All Permissions")

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

	logger.Log("Get Permission", permission.PermissionName)

	return permission, nil
}

func (s service) UpdatePermission(ctx context.Context, updateReq UpdatePermissionRequest) (string, error) {
	logger := log.With(s.logger, "method", "UpdatePermission")

	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

	updatePermission := db.UpdatePermissionParams{
		PermissionName: updateReq.PermissionName,
		Status:         updateReq.Status,
		UpdatedAt:      time,
		ID:             updateReq.ID,
	}

	err := s.repository.UpdatePermission(ctx, updatePermission)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Update Permission", updatePermission.PermissionName)

	return "Permission updated successfully", nil
}
