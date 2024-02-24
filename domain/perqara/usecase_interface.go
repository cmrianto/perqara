package perqara

import (
	"context"
	"perqara/domain/payload"
	"perqara/domain/request"
	"perqara/domain/response"
	"perqara/lib/helper"
)

type PerqaraUsecaseInterface interface {
	/* ------------------------	*/
	/* 			Auth			*/
	/* ------------------------ */
	// User Authentication
	Login(ctx context.Context, req *payload.LoginPayload) (*response.LoginResponse, error)
	ValidateAccess(ctx context.Context, req *helper.AuthenticationData) (bool, error)
	GetUserContentTypes(ctx context.Context) (*response.GetUserContentTypesResponse, error)

	// ContentTypes
	GetContentTypes(ctx context.Context, req *payload.GetContentTypesPayload) (*response.GetContentTypesResponse, error)
	CreateContentType(ctx context.Context, req *request.CreateContentTypeRequest) error

	// Permissions
	GetPermissionsDropdown(ctx context.Context, req *payload.GetPermissionsPayload) (*response.GetPermissionsResponse, error)

	// Roles
	GetRoles(ctx context.Context, req *payload.GetRolesPayload) (*response.GetRolesResponse, error)
	GetTotalRoles(ctx context.Context, req *payload.GetRolesPayload) (*response.GetTotalResponse, error)
	GetRoleDetail(ctx context.Context, req *payload.GetSingleIdPayload) (*response.Role, error)
	GetRolesDropdown(ctx context.Context, req *payload.GetRolesPayload) (*response.GetRolesResponse, error)
	CreateRole(ctx context.Context, req *request.CreateRoleRequest) error
	UpdateRole(ctx context.Context, req *request.UpdateRoleRequest) error
	DeleteRole(ctx context.Context, req *request.SingleIdRequest) error

	// Users
	GetUsers(ctx context.Context, req *payload.GetUsersPayload) (*response.GetUsersResponse, error)
	GetUserDetail(ctx context.Context, req *payload.GetSingleIdPayload) (*response.User, error)
	GetTotalUsers(ctx context.Context, req *payload.GetUsersPayload) (*response.GetTotalResponse, error)
	CreateUser(ctx context.Context, req *request.CreateUserRequest) error
	DeleteUser(ctx context.Context, req *request.SingleIdRequest) error
	/* ------------------------	*/
	/* 		End Of: Auth		*/
	/* ------------------------ */

	// Clients
	GetClients(ctx context.Context, req *payload.GetClientsPayload) (*response.GetClientsResponse, error)
	CreateClient(ctx context.Context, req *request.CreateClientRequest) error
	UpdateClient(ctx context.Context, req *request.UpdateClientRequest) error
	DeleteClient(ctx context.Context, req *request.SingleIdRequest) error
}
