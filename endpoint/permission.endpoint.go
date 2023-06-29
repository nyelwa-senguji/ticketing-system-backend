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
	CreatePermissionResponse struct {
		Status  int    `json:"status"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	UpdatePermissionResponse struct {
		Result string `json:"result"`
	}

	ListPermissionResponse struct {
		Result      string          `json:"result"`
		Permissions []db.Permission `json:"permissions"`
	}

	GetPermissionRequest struct {
		Id int32 `json:"id"`
	}

	GetPermissionResponse struct {
		Result     string        `json:"result"`
		Permission db.Permission `json:"permission"`
	}
)

func makeCreatePermissionEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		req := request.(service.CreatePermissionRequest)

		if reflect.DeepEqual(req.PermissionName, "") {
			return CreatePermissionResponse{Status: utils.StatusBadRequest, Success: false, Message: "Permission name can not be empty"}, nil
		}

		if reflect.DeepEqual(req.Status, "") {
			return CreatePermissionResponse{Status: utils.StatusBadRequest, Success: false, Message: "Permisison status can not be empty"}, nil
		}

		ok, err := s.CreatePermission(ctx, req)
		if err != nil {
			return CreatePermissionResponse{Status: utils.StatusInternalServerError, Success: false, Message: ok}, err
		}
		
		return CreatePermissionResponse{Status: utils.StatusOK, Success: true, Message: ok}, err
	}
}

func makeListPermissionEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ok, err := s.ListPermissions(ctx)
		return ListPermissionResponse{Result: "Permissions fetched Successfully", Permissions: ok}, err
	}
}

func makeGetPermissionEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetPermissionRequest)
		ok, err := s.GetPermission(ctx, req.Id)
		return GetPermissionResponse{Result: "Permission fetched Successfully", Permission: ok}, err
	}
}

func makeUpdatePermissionEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.UpdatePermissionRequest)
		ok, err := s.UpdatePermission(ctx, req)
		return UpdatePermissionResponse{Result: ok}, err
	}
}
