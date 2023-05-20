package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
)

type (
	CreatePermissionResponse struct {
		Result string `json:"result"`
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
		Result string `json:"result"`
		Permission db.Permission `json:"permission"`
	}
)

type Endpoint struct {
	CreatePermission endpoint.Endpoint
	ListPermission   endpoint.Endpoint
	GetPermission    endpoint.Endpoint
	UpdatePermission endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoint {
	return Endpoint{
		CreatePermission: makeCreatePermissionEndpoint(s),
		ListPermission:   makeListPermissionEndpoint(s),
		GetPermission:    makeGetPermission(s),
		UpdatePermission: makeUpdatePermission(s),
	}
}

func makeCreatePermissionEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.CreatePermissionRequest)
		ok, err := s.CreatePermission(ctx, req)
		return CreatePermissionResponse{Result: ok}, err
	}
}

func makeListPermissionEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ok, err := s.ListPermissions(ctx)
		return ListPermissionResponse{Result: "Permissions fetched Successfully", Permissions: ok}, err
	}
}

func makeGetPermission(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetPermissionRequest)
		ok, err := s.GetPermission(ctx, req.Id)
		return GetPermissionResponse{Result: "Permission fetched Successfully", Permission: ok}, err
	}
}

func makeUpdatePermission(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.UpdatePermissionRequest)
		ok, err := s.UpdatePermission(ctx, req)
		return UpdatePermissionResponse{Result: ok}, err
	}
}