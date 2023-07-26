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

func decodeCreateCategoryReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req service.CreateCategoryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeListCategoriesReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetCategoryReq(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	i, _ := strconv.ParseInt(vars["id"], 10, 32)
	id := int32(i)
	request := endpoint.GetCategoryRequest{
		Id: id,
	}
	return request, nil
}
