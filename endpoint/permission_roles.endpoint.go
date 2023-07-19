package endpoint

import (
	"context"
	"reflect"

	"github.com/go-kit/kit/endpoint"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
	"github.com/nyelwa-senguji/ticketing_system_backend/utils"
)

type (
	AssignPermissionToRoleResponse struct {
		Status  int    `json:"status"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	ListAssignedPermissionsToRoleResponse struct {
		Status       int     `json:"status"`
		Message      string  `json:"message"`
		PermissionID []int32 `json:"permission_id"`
	}

	AssignPermissionToRoleRequest struct {
		PermissionID int32 `json:"permission_id"`
		RoleID       int32 `json:"role_id"`
	}

	ListAssignedPermissionsToRoleRequest struct {
		RoleID int32 `json:"role_id"`
	}
)

func makeAssignPermissionToRole(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		req := request.(AssignPermissionToRoleRequest)

		if reflect.DeepEqual(req.PermissionID, nil) {
			return AssignPermissionToRoleResponse{Status: utils.StatusBadRequest, Success: true, Message: "PermissionID cannot be empty"}, nil
		}

		if reflect.DeepEqual(req.RoleID, nil) {
			return AssignPermissionToRoleResponse{Status: utils.StatusBadRequest, Success: true, Message: "RoleID cannot be empty"}, nil
		}

		ok, err := s.AssignPermissionToRole(ctx, req.RoleID, req.PermissionID)
		if err != nil {
			return AssignPermissionToRoleResponse{Status: utils.StatusBadRequest, Success: true, Message: err.Error()}, nil
		}

		return AssignPermissionToRoleResponse{Status: utils.StatusOK, Success: true, Message: ok}, err
	}
}

func makeListAssignedPermissionsToRole(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ListAssignedPermissionsToRoleRequest)
		ok, err := s.ListAssignedPermissionsToRole(ctx, req.RoleID)
		return ListAssignedPermissionsToRoleResponse{Status: utils.StatusOK, Message: "Permissions assigned to role", PermissionID: ok }, err
	}
}
