package transport

import (
	"context"
	"encoding/json"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/endpoint"
)

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware)

	r.Methods("POST").Path("/permissions").Handler(transport.NewServer(
		endpoints.CreatePermission,
		decodePermissionReq,
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
	var req db.CreatePermissionParams
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
