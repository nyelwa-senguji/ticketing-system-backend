package transport

import (
	"context"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/nyelwa-senguji/ticketing_system_backend/endpoint"
	"github.com/nyelwa-senguji/ticketing_system_backend/middleware"
	"github.com/nyelwa-senguji/ticketing_system_backend/utils"
)

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.NewMiddleware().HeaderMiddleware)

	a := r.PathPrefix("/").Subrouter()
	a.Use(middleware.NewMiddleware().AuthenticationMiddleware)
	
	/*****************************************************************
		Authentication transport layer
	******************************************************************/
	r.Methods("POST").Path("/login").Handler(transport.NewServer(
		endpoints.LoginUser,
		decodeLoginUserReq,
		utils.EncodeResponse,
	))

	/*****************************************************************
		Permissions transport layer
	******************************************************************/
	a.Methods("POST").Path("/permissions").Handler(transport.NewServer(
		endpoints.CreatePermission,
		decodeCreatePermissionReq,
		utils.EncodeResponse,
	))

	a.Methods("GET").Path("/permissions").Handler(transport.NewServer(
		endpoints.ListPermission,
		decodeListPermissionReq,
		utils.EncodeResponse,
	))

	a.Methods("GET").Path("/permissions/{id}").Handler(transport.NewServer(
		endpoints.GetPermission,
		decodeGetPermissionReq,
		utils.EncodeResponse,
	))

	a.Methods("PUT").Path("/permissions").Handler(transport.NewServer(
		endpoints.UpdatePermission,
		decodeUpdatePermissionReq,
		utils.EncodeResponse,
	))

	/*****************************************************************
		Roles transport layer
	******************************************************************/

	a.Methods("POST").Path("/roles").Handler(transport.NewServer(
		endpoints.CreateRole,
		decodeCreateRoleReq,
		utils.EncodeResponse,
	))

	a.Methods("GET").Path("/roles").Handler(transport.NewServer(
		endpoints.ListRoles,
		decodeListRolesReq,
		utils.EncodeResponse,
	))

	a.Methods("GET").Path("/roles/{id}").Handler(transport.NewServer(
		endpoints.GetRole,
		decodeGetRoleReq,
		utils.EncodeResponse,
	))

	a.Methods("PUT").Path("/roles").Handler(transport.NewServer(
		endpoints.UpdateRole,
		decodeUpdateRoleReq,
		utils.EncodeResponse,
	))

	/*****************************************************************
		Users transport layer
	******************************************************************/
	a.Methods("POST").Path("/users").Handler(transport.NewServer(
		endpoints.CreateUser,
		decodeCreateUserReq,
		utils.EncodeResponse,
	))

	return r
}

type request struct{}
