package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
)

type (
	CreateRoleResponse struct {
		Result string `json:"result"`
	}

	UpdateRoleResponse struct{
		Result string `json:"result"`
	}

	ListRolesResponse struct {
		Result string     `json:"result"`
		Roles  []db.Roles `json:"roles"`
	}

	GetRoleRequest struct {
		Id int32 `json:"id"`
	}

	GetRoleResponse struct {
		Result string   `json:"result"`
		Role   db.Roles `json:"role"`
	}
)

func makeCreateRoleEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.CreateRoleRequest)
		ok, err := s.CreateRole(ctx, req)
		return CreateRoleResponse{Result: ok}, err
	}
}

func makeListRolesEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ok, err := s.ListRoles(ctx)
		return ListRolesResponse{Result: "Roles fetched Successfully", Roles: ok}, err
	}
}

func makeGetRoleEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetRoleRequest)
		ok, err := s.GetRole(ctx, req.Id)
		return GetRoleResponse{Result: "Role fetched Successfully", Role: ok}, err
	}
}

func makeUpdateRoleEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.UpdateRoleRequest)
		ok, err := s.UpdateRole(ctx, req)
		return UpdateRoleResponse{Result: ok}, err
	}
}
