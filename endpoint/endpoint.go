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
	ListUsers  endpoint.Endpoint
	GetUser    endpoint.Endpoint

	/**********************************
		Permission Roles Endpoints
	***********************************/
	AssignPermissionToRole        endpoint.Endpoint
	ListAssignedPermissionsToRole endpoint.Endpoint
	RevokePermissionToRole        endpoint.Endpoint

	/**********************************
		Category Endpoints
	***********************************/
	CreateCategory endpoint.Endpoint
	GetCategory    endpoint.Endpoint
	ListCategories endpoint.Endpoint

	/**********************************
		Ticket Endpoints
	***********************************/
	CreateTicket endpoint.Endpoint
	ListTickets  endpoint.Endpoint

	/********************************
		Authentication Endpoints
	*********************************/
	LoginUser endpoint.Endpoint
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
		ListUsers:  makeListUsersEndpoint(s),
		GetUser:    makeGetUserEndpoint(s),

		AssignPermissionToRole:        makeAssignPermissionToRole(s),
		ListAssignedPermissionsToRole: makeListAssignedPermissionsToRole(s),
		RevokePermissionToRole:        makeRevokePermissionToRole(s),

		CreateCategory: makeCreateCategoryEndpoint(s),
		GetCategory:    makeGetCategoryEndpoint(s),
		ListCategories: makeListCategoriesEndpoint(s),

		CreateTicket: makeCreateTicketEndpoint(s),
		ListTickets: makeListTicketsEndpoint(s),

		LoginUser: makeLoginUserEndpoint(s),
	}
}
