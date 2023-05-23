package transport

import (
	"context"
	"encoding/json"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/nyelwa-senguji/ticketing_system_backend/endpoint"
)

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware)

	/*****************************************************************
		Permissions transport layer
	******************************************************************/
	r.Methods("POST").Path("/permissions").Handler(transport.NewServer(
		endpoints.CreatePermission,
		decodeCreatePermissionReq,
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

	r.Methods("PUT").Path("/permissions").Handler(transport.NewServer(
		endpoints.UpdatePermission,
		decodeUpdatePermissionReq,
		encodeResponse,
	))

	/*****************************************************************
		Roles transport layer
	******************************************************************/

	r.Methods("POST").Path("/roles").Handler(transport.NewServer(
		endpoints.CreateRole,
		decodeCreateRoleReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/roles").Handler(transport.NewServer(
		endpoints.ListRoles,
		decodeListRolesReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/roles/{id}").Handler(transport.NewServer(
		endpoints.GetRole,
		decodeGetRoleReq,
		encodeResponse,
	))

	r.Methods("PUT").Path("/roles").Handler(transport.NewServer(
		endpoints.UpdateRole,
		decodeUpdateRoleReq,
		encodeResponse,
	))

	return r
}

type request struct{}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
