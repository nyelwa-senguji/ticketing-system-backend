package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/utils"
)

func (s service) AssignPermissionToRole(ctx context.Context, roleID int32, permissionID int32) (string, error) {

	logger := log.With(s.logger, "method", "Assigning Permission to role")

	assignedPermissions, err := s.repository.ListAssignedPermissionsToRole(ctx, roleID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	checkIfExists := utils.IsAvailable(assignedPermissions, permissionID)
	if checkIfExists {
		return "This Permission is already assigned to this role", nil
	}

	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

	assignPermissionToRole := db.AssignPermissionToRoleParams{
		RoleID:       roleID,
		PermissionID: permissionID,
		UpdatedAt: time,
		CreatedAt: time,
	}

	_, err = s.repository.AssignPermissionToRole(ctx, assignPermissionToRole)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Assigning Permission to role")

	return "Permission assigned successfully", nil
}

func (s service) ListAssignedPermissionsToRole(ctx context.Context, roleID int32) ([]int32, error) {

	logger := log.With(s.logger, "method", "Listing Assigned Permission to Role")

	permissions, err := s.repository.ListAssignedPermissionsToRole(ctx, roleID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	logger.Log("List All Permissions Assigned to role")

	return permissions, nil
}

func (s service) RevokePermissionToRole(ctx context.Context, roleID int32, permissionID int32) (string, error){
	logger := log.With(s.logger, "method", "Listing Assigned Permission to Role")

	revokePermissionRole := db.RevokePermissionRoleParams{
		RoleID: roleID,
		PermissionID: permissionID,
	}

	_, err := s.repository.RevokePermissionRole(ctx, revokePermissionRole)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Permissions revoked successfully to role")
	
	return "Permissions revoked successfully to role", nil
}
