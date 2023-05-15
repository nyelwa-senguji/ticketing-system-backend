package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
)

type (
	CreatePermissionResponse struct{
		Ok string `json:"ok"`
	}
)

type Endpoint struct {
	CreatePermission endpoint.Endpoint
	ListPermissions  endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoint {
	return Endpoint{
		CreatePermission: makeCreatePermissionEndpoint(s),
	}
}

func makeCreatePermissionEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(db.CreatePermissionParams)
		ok, err := s.CreatePermission(ctx, req)
		return CreatePermissionResponse{Ok: ok}, err
	}
}


