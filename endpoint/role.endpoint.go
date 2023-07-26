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
	CreateRoleResponse struct {
		Status  int    `json:"status"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	UpdateRoleResponse struct {
		Status  int    `json:"status"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	ListRolesResponse struct {
		Status      int             `json:"status"`
		Message     string          `json:"message"`
		Roles  []db.Roles `json:"roles"`
	}

	GetRoleRequest struct {
		Id int32 `json:"id"`
	}

	GetRoleResponse struct {
		Status      int             `json:"status"`
		Message     string          `json:"message"`
		Role   db.Roles `json:"role"`
	}
)

func makeCreateRoleEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		req := request.(service.CreateRoleRequest)

		if reflect.DeepEqual(req.RoleName, "") {
			return CreateRoleResponse{Status: utils.StatusBadRequest, Success: false, Message: "Role name cannot be empty"}, nil
		}

		if reflect.DeepEqual(req.Status, "") {
			return CreateRoleResponse{Status: utils.StatusBadRequest, Success: false, Message: "Role status cannot be empty"}, nil
		}
		ok, err := s.CreateRole(ctx, req)
		if err != nil {
			return CreateRoleResponse{Status: utils.StatusBadRequest, Success: false, Message: err.Error()}, nil
		}
		return CreateRoleResponse{Status: utils.StatusOK, Success: true, Message: ok}, err
	}
}

func makeListRolesEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ok, err := s.ListRoles(ctx)
		return ListRolesResponse{Status: utils.StatusOK, Message: "Roles fetched Successfully", Roles: ok}, err
	}
}

func makeGetRoleEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRoleRequest)
		ok, err := s.GetRole(ctx, req.Id)
		if err != nil {
			return GetRoleResponse{Status: utils.StatusBadRequest, Message: err.Error(), Role: ok}, nil
		}
		return GetRoleResponse{Status: utils.StatusOK, Message: "Role fetched Successfuly", Role: ok}, err
	}
}

func makeUpdateRoleEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.UpdateRoleRequest)
		ok, err := s.UpdateRole(ctx, req)
		return UpdateRoleResponse{Status: utils.StatusOK, Success: true, Message: ok}, err
	}
}
