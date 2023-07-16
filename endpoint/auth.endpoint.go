package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
	"github.com/nyelwa-senguji/ticketing_system_backend/utils"
)

type (
	LoginUserResponse struct {
		Status  int                          `json:"status"`
		Message string                       `json:"message"`
		User    *service.LoggedInUserResponse `json:"user"`
	}
)

func makeLoginUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(db.LoginUserParams)
		ok, _, err := s.LoginUser(ctx, req)
		return LoginUserResponse{Status: utils.StatusOK, Message: "User logged in successfully", User: ok}, err
	}
}
