package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/nyelwa-senguji/ticketing_system_backend/service"
)

type Endpoint struct {
	/************************
		Permission Endpoints
	*************************/
	CreatePermission endpoint.Endpoint
	ListPermission   endpoint.Endpoint
	GetPermission    endpoint.Endpoint
	UpdatePermission endpoint.Endpoint

	/************************
		Role Endpoints
	*************************/
	CreateRole endpoint.Endpoint
	ListRoles  endpoint.Endpoint
	GetRole    endpoint.Endpoint
	UpdateRole endpoint.Endpoint

	/************************
		User Endpoints
	*************************/
	CreateUser endpoint.Endpoint
	ListUsers endpoint.Endpoint
	GetUser endpoint.Endpoint

	/**********************************
		Permission Roles Endpoints
	***********************************/
	AssignPermissionToRole endpoint.Endpoint
	ListAssignedPermissionsToRole endpoint.Endpoint

	/********************************
		Authentication Endpoints
	*********************************/
	LoginUser  endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoint {
	return Endpoint{
		CreatePermission: makeCreatePermissionEndpoint(s),
		ListPermission:   makeListPermissionEndpoint(s),
		GetPermission:    makeGetPermissionEndpoint(s),
		UpdatePermission: makeUpdatePermissionEndpoint(s),

		CreateRole: makeCreateRoleEndpoint(s),
		ListRoles:  makeListRolesEndpoint(s),
		GetRole:    makeGetRoleEndpoint(s),
		UpdateRole: makeUpdateRoleEndpoint(s),

		CreateUser: makeCreateUserEndpoint(s),
		ListUsers: makeListUsersEndpoint(s),
		GetUser: makeGetUserEndpoint(s),

		AssignPermissionToRole: makeAssignPermissionToRole(s),
		
		LoginUser:  makeLoginUserEndpoint(s),
	}
}
