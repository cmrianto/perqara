package perqara

import (
	"context"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"

	"gorm.io/gorm"
)

type PerqaraRepoInterface interface {
	/* ------------------------	*/
	/* 			Auth			*/
	/* ------------------------ */
	// ContentTypes
	GetContentTypes(ctx context.Context, req *payload.GetContentTypesPayload) (*[]model.ContentType, error)
	CreateContentType(ctx context.Context, tx *gorm.DB, data *model.ContentType) error

	// Permissions
	GetPermissions(ctx context.Context, req *payload.GetPermissionsPayload) (*[]model.Permission, error)
	CreatePermissions(ctx context.Context, tx *gorm.DB, data *[]model.Permission) error

	// Roles
	GetRoles(ctx context.Context, req *payload.GetRolesPayload) ([]*model.Role, error)
	GetTotalRoles(ctx context.Context, req *payload.GetRolesPayload) (int64, error)
	CreateRole(ctx context.Context, tx *gorm.DB, data *model.Role) error
	UpdateRole(ctx context.Context, tx *gorm.DB, req *payload.UpdateRolePayload) error
	DeleteRoles(ctx context.Context, tx *gorm.DB, req *payload.DeleteRolesPayload) error

	// RolePermissions
	GetRolePermissions(ctx context.Context, req *payload.GetRolePermissionsPayload) ([]*model.RolePermission, error)
	CreateRolePermissions(ctx context.Context, tx *gorm.DB, data *[]model.RolePermission) error
	DeleteRolePermissions(ctx context.Context, tx *gorm.DB, req *payload.DeleteRolePermissionsPayload) error

	// Users
	GetUsers(ctx context.Context, req *payload.GetUsersPayload) ([]*model.User, error)
	GetTotalUsers(ctx context.Context, req *payload.GetUsersPayload) (int64, error)
	CreateUser(ctx context.Context, tx *gorm.DB, data *model.User) error
	DeleteUsers(ctx context.Context, tx *gorm.DB, req *payload.DeleteUsersPayload) error

	// User Roles
	GetUserRoles(ctx context.Context, req *payload.GetUserRolesPayload) ([]*model.UserRole, error)
	CreateUserRoles(ctx context.Context, tx *gorm.DB, data []*model.UserRole) error
	DeleteUserRoles(ctx context.Context, tx *gorm.DB, req *payload.DeleteUserRolesPayload) error
	/* ------------------------	*/
	/* 		End Of: Auth		*/
	/* ------------------------ */

	// Foreign Key Switch
	ForeignKeySwitch(in int)

	// Clients
	GetClients(ctx context.Context, req *payload.GetClientsPayload) (*[]model.Client, error)
	CreateClient(ctx context.Context, tx *gorm.DB, data *model.Client) error
	UpdateClient(ctx context.Context, tx *gorm.DB, req *payload.UpdateClientPayload) error
	DeleteClients(ctx context.Context, tx *gorm.DB, req *payload.DeleteClientsPayload) error
}
