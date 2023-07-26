package service

import (
	"context"
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

	logger := log.With(s.logger, "method", "CreateRole")

	checkIfRoleExists, _ := s.repository.GetRoleByName(ctx, createRoleReq.RoleName)

	if checkIfRoleExists.RoleName == createRoleReq.RoleName {
		return "Role already exists", nil
	}

	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

	role := db.CreateRoleParams{
		RoleName:  createRoleReq.RoleName,
		Status:    createRoleReq.Status,
		UpdatedAt: time,
		CreatedAt: time,
	}

	_, err := s.repository.CreateRole(ctx, role)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Create Role", role.RoleName)

	return "Role Created Successfully", nil
}

func (s service) ListRoles(ctx context.Context) ([]db.Roles, error) {
	logger := log.With(s.logger, "method", "ListRoles")

	roles, err := s.repository.ListRoles(ctx)

	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	logger.Log("List All Roles")

	return roles, nil
}

func (s service) GetRole(ctx context.Context, id int32) (db.Roles, error) {
	logger := log.With(s.logger, "method", "GetRole")

	role, err := s.repository.GetRole(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return role, err
	}

	logger.Log("Get Role", role.RoleName)

	return role, nil
}

func (s service) UpdateRole(ctx context.Context, updateRoleReq UpdateRoleRequest) (string, error) {
	logger := log.With(s.logger, "method", "UpdatePermission")

	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

	updateRole := db.UpdateRoleParams{
		RoleName:  updateRoleReq.RoleName,
		Status:    updateRoleReq.Status,
		UpdatedAt: time,
		ID:        updateRoleReq.ID,
	}

	err := s.repository.UpdateRole(ctx, updateRole)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Update Role", updateRole.RoleName)

	return "Role Updated Successfully", nil
}
