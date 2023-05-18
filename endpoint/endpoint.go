package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
)

type (
	CreatePermissionResponse struct{
		Result string `json:"result"`
	}

	ListPermissionResponse struct{
		Result string `json:"result"`
		Permissions []db.Permission `json:"permissions"`
	}
)

type Endpoint struct {
	CreatePermission endpoint.Endpoint
	ListPermission  endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoint {
	return Endpoint{
		CreatePermission: makeCreatePermissionEndpoint(s),
		ListPermission: makeCreateListPermissionEndpoint(s),
	}
}

func makeCreatePermissionEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(db.CreatePermissionParams)
		ok, err := s.CreatePermission(ctx, req)
		return CreatePermissionResponse{Result: ok}, err
	}
}

func makeCreateListPermissionEndpoint(s service.Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error){
		ok, err := s.ListPermissions(ctx)
		return ListPermissionResponse{Result: "Permissions fetched Successfully", Permissions: ok}, err
	}
}


