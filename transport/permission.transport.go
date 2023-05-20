package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/nyelwa-senguji/ticketing_system_backend/endpoint"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
)

type request struct{}

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware)

	r.Methods("POST").Path("/permissions").Handler(transport.NewServer(
		endpoints.CreatePermission,
		decodePermissionReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/permissions").Handler(transport.NewServer(
		endpoints.ListPermission,
		decodeListPermissionReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/permissions/{id}").Handler(transport.NewServer(
		endpoints.GetPermission,
		decodeGetPermissionReq,
		encodeResponse,
	))

	return r
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodePermissionReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req service.CreatePermissionRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeListPermissionReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetPermissionReq(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	i, _ := strconv.ParseInt(vars["id"], 10, 32)
	id := int32(i)
	request := endpoint.GetPermissionRequest{
		Id: id,
	}
	return request, nil
}
