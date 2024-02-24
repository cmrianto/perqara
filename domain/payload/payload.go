package payload

import (
	"perqara/lib/helper"
)

type GetServicesPayload struct {
	ServiceIds    []int64
	Search        *string
	Limit         *int32
	Offset        *int32
	LastCreatedAt *string
	AuthData      helper.AuthenticationData
}

type GetPermissionsPayload struct {
	PermissionIds   []int64
	ContentTypeIds  []int64
	Search          *string
	PermissionNames []string
	AuthData        helper.AuthenticationData
}

type GetRolePermissionsPayload struct {
	RolePermissionIds []int64
	RoleIds           []int64
}

type DeleteUserRolesPayload struct {
	UserIds []int64
}

type GetRolesPayload struct {
	RoleIds       []int64
	Search        *string
	Limit         *int32
	Offset        *int32
	LastCreatedAt *string
	AuthData      helper.AuthenticationData
}

type DeleteUsersPayload struct {
	UserIds []int64
}

type DeleteRolesPayload struct {
	RoleIds []int64
}

type DeleteRolePermissionsPayload struct {
	RoleIds []int64
}

type GetUserRolesPayload struct {
	UserIds *[]int64
}

type GetUsersPayload struct {
	UserIds         []int64
	Username        *string
	FilterUserTypes []int64
	Search          *string
	Limit           *int32
	Offset          *int32
	LastCreatedAt   *string
	AuthData        helper.AuthenticationData
}

type UpdateRolePayload struct {
	Id   *int64
	Name *string
}

type LoginPayload struct {
	Username string
	Password string
}

type GetContentTypesPayload struct {
	ContentTypeIds []int64
	Search         *string
	Limit          *int32
	Offset         *int32
	LastCreatedAt  *string
}

type GetSingleIdPayload struct {
	Id       *int64
	AuthData helper.AuthenticationData
}

type GetClientsPayload struct {
	ClientIds     *[]int64
	Search        *string
	Limit         *int32
	Offset        *int32
	LastCreatedAt *string
	AuthData      helper.AuthenticationData
}

type UpdateClientPayload struct {
	ClientId int64
	Name     string
}

type DeleteClientsPayload struct {
	ClientIds []int64
}
