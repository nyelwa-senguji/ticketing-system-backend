package service

import (
	"context"
	"reflect"
	"strconv"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/utils"
)

func (s service) LoginUser(ctx context.Context, loginUserReq db.LoginUserParams) (string, error) {
	logger := log.With(s.logger, "method", "LoginUser")

	user := db.LoginUserParams{
		Username: loginUserReq.Username,
		Password: loginUserReq.Password,
	}

	if reflect.DeepEqual(user.Username, "") {
		return "Username cannot be empty", nil
	}

	if reflect.DeepEqual(user.Password, "") {
		return "Password cannot be empty", nil
	}

	loggedInUser, err := s.repository.LoginUser(ctx, user)
	if err != nil {
		return "Incorrect Username/Password", level.Error(logger).Log("err", err)
	}

	token := utils.GenerateJWTToken(strconv.Itoa(int(loggedInUser.ID)))

	logger.Log("User Logged In Successfuly", user.Username)

	return token, nil
}
