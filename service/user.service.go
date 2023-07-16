package service

import (
	"context"
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

	_, err := s.repository.CreateUser(ctx, user)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Create User", user.Username)

	return "User Created Successfully", nil
}

func (s service) ListUsers(ctx context.Context) ([]db.Users, error){
	logger := log.With(s.logger, "method", "ListUsers")

	users, err := s.repository.ListUsers(ctx)

	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	logger.Log("List All Users")

	return users, nil
}

func (s service) GetUser(ctx context.Context, id int32) (db.Users, error){
	logger := log.With(s.logger, "method", "Getuser")

	user, err := s.repository.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return user, err
	}

	logger.Log("Get User", user.Username)

	return user, nil
}
