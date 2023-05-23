package service

import (
	"context"
	"reflect"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
)

type CreateRoleRequest struct {
	RoleName string `json:"role_name"`
	Status   string `json:"status"`
}

type UpdateRoleRequest struct {
	RoleName string `json:"role_name"`
	Status   string `json:"status"`
	ID       int32  `json:"id"`
}

func (s service) CreateRole(ctx context.Context, createRoleReq CreateRoleRequest) (string, error) {

	logger := log.With(s.logger, "method", "CreatePermission")

	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

	role := db.CreateRoleParams{
		RoleName:  createRoleReq.RoleName,
		Status:    createRoleReq.Status,
		UpdatedAt: time,
		CreatedAt: time,
	}

	if reflect.DeepEqual(role.RoleName, "") {
		return "Role name cannot be empty", nil
	}

	if reflect.DeepEqual(role.Status, "") {
		return "Role status cannot be empty", nil
	}

	_, err := s.repository.CreateRole(ctx, role)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Create Permission", role.RoleName)

	return "Role Created Successfully", nil
}

func (s service) ListRoles(ctx context.Context) ([]db.Roles, error) {
	return nil, nil
}

func (s service) GetRole(ctx context.Context, id int32) (db.Roles, error) {
	return db.Roles{}, nil
}

func (s service) UpdateRole(ctx context.Context, updateRoleReq UpdateRoleRequest) (string, error) {
	return "", nil
}
