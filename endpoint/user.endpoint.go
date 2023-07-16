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
	CreateUserResponse struct {
		Status  int    `json:"status"`
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	ListUsersResponse struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Users []db.Users `json:"users"`
	}

	GetUserRequest struct {
		Id int32 `json:"id"`
	}

	GetUserResponse struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		User db.Users `json:"user"`
	}
)

func makeCreateUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(service.CreateUserRequest)

		if reflect.DeepEqual(req.Username, "") {
			return CreateUserResponse{Status: utils.StatusBadRequest, Success: false, Message: "Username cannot be empty"}, nil
		}

		if reflect.DeepEqual(req.Password, "") {
			return CreateUserResponse{Status: utils.StatusBadRequest, Success: false, Message: "Password cannot be empty"}, nil
		}

		if reflect.DeepEqual(req.RoleID, "") {
			return CreateUserResponse{Status: utils.StatusBadRequest, Success: false, Message: "RoleID cannot be empty"}, nil
		}

		ok, err := s.CreateUser(ctx, req)

		if err != nil {
			return CreateUserResponse{Status: utils.StatusBadRequest, Success: false, Message: err.Error()}, nil
		}

		return CreateUserResponse{Status: utils.StatusOK, Success: true, Message: ok}, err
	}
}

func makeListUsersEndpoint(s service.Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		ok, err := s.ListUsers(ctx)
		return ListUsersResponse{Status: utils.StatusOK, Message: "Users fetched Successfully", Users: ok}, err
	}
}

func makeGetUserEndpoint(s service.Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetUserRequest)
		ok, err := s.GetUser(ctx, req.Id)
		if err != nil {
			return GetUserResponse{Status: utils.StatusBadRequest, Message: err.Error(), User: ok}, nil
		}
		return GetUserResponse{Status: utils.StatusOK, Message: "User fetched successfully", User: ok}, err
	}
}
