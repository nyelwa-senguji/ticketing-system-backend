package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
)

type (
	CreateUserResponse struct {
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
