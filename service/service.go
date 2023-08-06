package service

import (
	"context"

	"github.com/go-kit/kit/log"

	db "github.com/nyelwa-senguji/ticketing_system_backend/db/sqlc"
	"github.com/nyelwa-senguji/ticketing_system_backend/token"
	"github.com/nyelwa-senguji/ticketing_system_backend/utils"
)

type Service interface {
	/******************************
		Services for Permission
	*******************************/
	CreatePermission(ctx context.Context, createPermissionParams CreatePermissionRequest) (string, error)
	ListPermissions(ctx context.Context) ([]db.Permission, error)
	GetPermission(ctx context.Context, id int32) (db.Permission, error)
	UpdatePermission(ctx context.Context, updateReq UpdatePermissionRequest) (string, error)

	/******************************
		Services for Roles
	*******************************/
	CreateRole(ctx context.Context, createRoleReq CreateRoleRequest) (string, error)
	ListRoles(ctx context.Context) ([]db.Roles, error)
	GetRole(ctx context.Context, id int32) (db.Roles, error)
	UpdateRole(ctx context.Context, updateRoleReq UpdateRoleRequest) (string, error)

	/******************************
		Services for Users
	*******************************/
	CreateUser(ctx context.Context, createUserRequest CreateUserRequest) (string, error)
	ListUsers(ctx context.Context) ([]db.Users, error)
	GetUser(ctx context.Context, id int32) (db.Users, error)

	/***********************************
		Services for Permission Roles
	************************************/
	AssignPermissionToRole(ctx context.Context, roleID int32, permissionID int32) (string, error)
	ListAssignedPermissionsToRole(ctx context.Context, roleID int32) ([]int32, error)
	RevokePermissionToRole(ctx context.Context, roleID int32, permissionID int32) (string, error)

	/***********************************
		Services for Category
	************************************/
	CreateCategory(ctx context.Context, createCategoryReq CreateCategoryRequest) (string, error)
	GetCategory(ctx context.Context, id int32)(db.Category, error)
	ListCategories(ctx context.Context) ([]db.Category, error)

	/***********************************
		Services for Tickets
	************************************/
	CreateTicket(ctx context.Context, createTicketReq CreateTicketRequest) (string, error)
	ListTickets(ctx context.Context) ([]ListTicketsResponse, error)

	/******************************
		Services for Authentication
	*******************************/
	LoginUser(ctx context.Context, loginUserReq db.LoginUserParams) (*LoggedInUserResponse, string, error)
}

type service struct {
	repository *db.Repository
	logger     log.Logger
	tokenMaker token.Maker
}

func NewService(repo *db.Repository, logger log.Logger) Service {
	tokenMaker, _ := token.NewPasetoMaker(utils.LoadEnviromentalVariables("SECRET_KEY"))
	return &service{
		repository: repo,
		logger:     logger,
		tokenMaker: tokenMaker,
	}
}
