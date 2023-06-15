package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
)

type (
	CreateUserResponse struct {
		Result string `json:"result"`
	}
	LoginUserResponse struct {
		Result string `json:"result"`
	}
)

func makeCreateUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.CreateUserRequest)
		ok, err := s.CreateUser(ctx, req)
		return CreateUserResponse{Result: ok}, err
	}
}

func makeLoginUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(db.LoginUserParams)
		ok, err := s.LoginUser(ctx, req)
		return LoginUserResponse{Result: ok}, err
	}
}
