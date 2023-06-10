package service

import (
	"context"
	"reflect"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
)

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   int32  `json:"role_id"`
}

func (s service) CreateUser(ctx context.Context, createUserReq CreateUserRequest) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	time, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))

	user := db.CreateUserParams{
		Username:  createUserReq.Username,
		Password:  createUserReq.Password,
		CreatedAt: time,
		UpdatedAt: time,
		RoleID:    createUserReq.RoleID,
	}

	if reflect.DeepEqual(user.Username, "") {
		return "Username cannot be empty", nil
	}

	if reflect.DeepEqual(user.Password, "") {
		return "Password cannot be empty", nil
	}

	if reflect.DeepEqual(user.RoleID, nil) {
		return "RoleID cannot be empty", nil
	}

	_, err := s.repository.CreateUser(ctx, user)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Create User", user.Username)

	return "User Created Successfully", nil
}
