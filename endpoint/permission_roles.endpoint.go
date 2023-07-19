package endpoint

import (
	"context"
	"reflect"

	"github.com/go-kit/kit/endpoint"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
	"github.com/nyelwa-senguji/ticketing_system_backend/utils"
)

type (
	AssignPermissionToRoleResponse struct {
		Status  int    `json:"status"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

func makeAssignPermissionToRole(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		req := request.(db.AssignPermissionToRoleParams)

		if reflect.DeepEqual(req.PermissionID, nil){
			return AssignPermissionToRoleResponse{Status: utils.StatusBadRequest, Success: true, Message: "PermissionID cannot be empty"}, nil
		}

		if reflect.DeepEqual(req.RoleID, nil){
			return AssignPermissionToRoleResponse{Status: utils.StatusBadRequest, Success: true, Message: "RoleID cannot be empty"}, nil
		}

		ok, err := s.AssignPermissionToRole(ctx, req.RoleID, req.PermissionID)
		if err != nil {
			return AssignPermissionToRoleResponse{Status: utils.StatusBadRequest, Success: true, Message: err.Error()}, nil
		}

		return AssignPermissionToRoleResponse{Status: utils.StatusBadRequest, Success: true, Message: ok}, err
	}
}
