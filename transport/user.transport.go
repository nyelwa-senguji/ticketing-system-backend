package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nyelwa-senguji/ticketing_system_backend/endpoint"
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

func decodeListUsersReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	i, _ := strconv.ParseInt(vars["id"], 10, 32)
	id := int32(i)
	request := endpoint.GetUserRequest{
		Id: id,
	}
	return request, nil
}
