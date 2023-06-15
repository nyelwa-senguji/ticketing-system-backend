package transport

import (
	"context"
	"encoding/json"
	"net/http"

	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
)

func decodeCreateUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req service.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeLoginUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req db.LoginUserParams
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
