package service

import (
	"context"
	"reflect"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
)

type LoggedInUserResponse struct {
	UserID int32 `json:"user_id"`
	RoleName string `json:"role_name"`
	Token    string `json:"token"`
	Username string `json:"username"`
}

func (s service) LoginUser(ctx context.Context, loginUserReq db.LoginUserParams) (*LoggedInUserResponse, string, error) {
	logger := log.With(s.logger, "method", "LoginUser")

	user := db.LoginUserParams{
		Username: loginUserReq.Username,
		Password: loginUserReq.Password,
	}

	if reflect.DeepEqual(user.Username, "") {
		return nil, "Username cannot be empty", nil
	}

	if reflect.DeepEqual(user.Password, "") {
		return nil, "Password cannot be empty", nil
	}

	userRow, err := s.repository.LoginUser(ctx, user)
	if err != nil {
		return nil, "Incorrect Username/Password", level.Error(logger).Log("err", err)
	}

	token, err := s.tokenMaker.CreateToken(user.Username, 30 * time.Minute)
	if err != nil {
		return nil, "Could not create token", level.Error(logger).Log("err", err)
	}

	logger.Log("User Logged In Successfuly", user.Username)

	role, _ := s.repository.GetRole(ctx, userRow.RoleID)

	loggedUser := LoggedInUserResponse{
		UserID: userRow.ID,
		RoleName: role.RoleName,
		Username: user.Username,
		Token: token,
	}

	return  &loggedUser, "", nil
}
