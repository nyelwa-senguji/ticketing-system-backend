package transport

import (
	"context"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/nyelwa-senguji/ticketing_system_backend/endpoint"
	"github.com/nyelwa-senguji/ticketing_system_backend/utils"
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
		utils.EncodeResponse,
	))

	r.Methods("GET").Path("/permissions").Handler(transport.NewServer(
		endpoints.ListPermission,
		decodeListPermissionReq,
		utils.EncodeResponse,
	))

	r.Methods("GET").Path("/permissions/{id}").Handler(transport.NewServer(
		endpoints.GetPermission,
		decodeGetPermissionReq,
		utils.EncodeResponse,
	))

	r.Methods("PUT").Path("/permissions").Handler(transport.NewServer(
		endpoints.UpdatePermission,
		decodeUpdatePermissionReq,
		utils.EncodeResponse,
	))

	/*****************************************************************
		Roles transport layer
	******************************************************************/

	r.Methods("POST").Path("/roles").Handler(transport.NewServer(
		endpoints.CreateRole,
		decodeCreateRoleReq,
		utils.EncodeResponse,
	))

	r.Methods("GET").Path("/roles").Handler(transport.NewServer(
		endpoints.ListRoles,
		decodeListRolesReq,
		utils.EncodeResponse,
	))

	r.Methods("GET").Path("/roles/{id}").Handler(transport.NewServer(
		endpoints.GetRole,
		decodeGetRoleReq,
		utils.EncodeResponse,
	))

	r.Methods("PUT").Path("/roles").Handler(transport.NewServer(
		endpoints.UpdateRole,
		decodeUpdateRoleReq,
		utils.EncodeResponse,
	))

	/*****************************************************************
		Users transport layer
	******************************************************************/
	r.Methods("POST").Path("/users").Handler(transport.NewServer(
		endpoints.CreateUser,
		decodeCreateUserReq,
		utils.EncodeResponse,
	))

	r.Methods("POST").Path("/login").Handler(transport.NewServer(
		endpoints.LoginUser,
		decodeLoginUserReq,
		utils.EncodeResponse,
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


