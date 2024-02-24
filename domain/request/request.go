package request

import (
	"perqara/domain/constant"
	"perqara/domain/payload"
	"perqara/lib/helper"
	"strconv"
)

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (o *LoginRequest) ToPayload() *payload.LoginPayload {
	if o == nil {
		return nil
	}
	return &payload.LoginPayload{
		Username: o.UserName,
		Password: o.Password,
	}
}

type GetContentTypesRequest struct {
	Search        string `json:"search"`
	Limit         string `json:"size"`
	Offset        string `json:"page"`
	LastCreatedAt string `json:"last_created_at"`
}

func (o *GetContentTypesRequest) ToPayload() *payload.GetContentTypesPayload {
	if o == nil {
		return nil
	}
	limitVal, _ := strconv.ParseInt(o.Limit, 10, 32)
	limit := int32(limitVal)
	offsetVal, _ := strconv.ParseInt(o.Offset, 10, 32)
	offset := int32(offsetVal)
	offset = (offset - 1) * limit
	return &payload.GetContentTypesPayload{
		Search:        &o.Search,
		Limit:         &limit,
		Offset:        &offset,
		LastCreatedAt: &o.LastCreatedAt,
	}
}

type GetUsersRequest struct {
	FilterUserTypes []string                  `json:"filter_user_types"`
	Search          string                    `json:"search"`
	Limit           string                    `json:"size"`
	Offset          string                    `json:"page"`
	LastCreatedAt   string                    `json:"last_created_at"`
	AuthData        helper.AuthenticationData `swaggerignore:"true"`
}

func (o *GetUsersRequest) ToPayload() *payload.GetUsersPayload {
	if o == nil {
		return nil
	}
	limitVal, _ := strconv.ParseInt(o.Limit, 10, 32)
	limit := int32(limitVal)
	offsetVal, _ := strconv.ParseInt(o.Offset, 10, 32)
	offset := int32(offsetVal)
	offset = (offset - 1) * limit
	filterUserTypes := []int64{}
	for _, fut := range o.FilterUserTypes {
		userType, exist := constant.UserTypeValue[fut]
		if exist {
			filterUserTypes = append(filterUserTypes, userType)
		}
	}
	return &payload.GetUsersPayload{
		FilterUserTypes: filterUserTypes,
		Search:          &o.Search,
		Limit:           &limit,
		Offset:          &offset,
		LastCreatedAt:   &o.LastCreatedAt,
		AuthData:        o.AuthData,
	}
}

type GetClientsRequest struct {
	FilterProvinceId          string                    `json:"client_province_id"`
	FilterWarehouseProvinceId string                    `json:"warehouse_province_id"`
	Search                    string                    `json:"search"`
	Limit                     string                    `json:"size"`
	Offset                    string                    `json:"page"`
	LastCreatedAt             string                    `json:"last_created_at"`
	AuthData                  helper.AuthenticationData `swaggerignore:"true"`
}

func (o *GetClientsRequest) ToPayload() *payload.GetClientsPayload {
	if o == nil {
		return nil
	}
	limitVal, _ := strconv.ParseInt(o.Limit, 10, 32)
	limit := int32(limitVal)
	offsetVal, _ := strconv.ParseInt(o.Offset, 10, 32)
	offset := int32(offsetVal)
	offset = (offset - 1) * limit
	provinceIdVal := []int64{}
	provinceId, _ := strconv.ParseInt(o.FilterProvinceId, 10, 64)
	if provinceId > 0 {
		provinceIdVal = append(provinceIdVal, provinceId)
	}
	filterWarehouseProvinceIdVal := []int64{}
	filterWarehouseProvinceId, _ := strconv.ParseInt(o.FilterWarehouseProvinceId, 10, 64)
	if filterWarehouseProvinceId > 0 {
		filterWarehouseProvinceIdVal = append(filterWarehouseProvinceIdVal, filterWarehouseProvinceId)
	}
	return &payload.GetClientsPayload{
		Search:        &o.Search,
		Limit:         &limit,
		Offset:        &offset,
		LastCreatedAt: &o.LastCreatedAt,
		AuthData:      o.AuthData,
	}
}

type GetRolesDropdownRequest struct {
	Search   string                    `json:"search"`
	AuthData helper.AuthenticationData `swaggerignore:"true"`
}

func (o *GetRolesDropdownRequest) ToPayload() *payload.GetRolesPayload {
	if o == nil {
		return nil
	}
	return &payload.GetRolesPayload{
		Search:   &o.Search,
		AuthData: o.AuthData,
	}
}

type GetRolesRequest struct {
	Search        string                    `json:"search"`
	Limit         string                    `json:"size"`
	Offset        string                    `json:"page"`
	LastCreatedAt string                    `json:"last_created_at"`
	AuthData      helper.AuthenticationData `swaggerignore:"true"`
}

func (o *GetRolesRequest) ToPayload() *payload.GetRolesPayload {
	if o == nil {
		return nil
	}
	limitVal, _ := strconv.ParseInt(o.Limit, 10, 32)
	limit := int32(limitVal)
	offsetVal, _ := strconv.ParseInt(o.Offset, 10, 32)
	offset := int32(offsetVal)
	offset = (offset - 1) * limit
	return &payload.GetRolesPayload{
		Search:        &o.Search,
		Limit:         &limit,
		Offset:        &offset,
		LastCreatedAt: &o.LastCreatedAt,
		AuthData:      o.AuthData,
	}
}

type GetPermissionsRequest struct {
	Search   string                    `json:"search"`
	AuthData helper.AuthenticationData `swaggerignore:"true"`
}

func (o *GetPermissionsRequest) ToPayload() *payload.GetPermissionsPayload {
	if o == nil {
		return nil
	}
	return &payload.GetPermissionsPayload{
		Search:   &o.Search,
		AuthData: o.AuthData,
	}
}

type GetSingleIdRequest struct {
	Id       string                    `json:"id"`
	AuthData helper.AuthenticationData `swaggerignore:"true"`
}

func (o *GetSingleIdRequest) ToPayload() *payload.GetSingleIdPayload {
	if o == nil {
		return nil
	}
	idVal, _ := strconv.ParseInt(o.Id, 10, 64)
	return &payload.GetSingleIdPayload{
		Id:       &idVal,
		AuthData: o.AuthData,
	}
}

type SingleIdRequest struct {
	Id       string                    `json:"id"`
	AuthData helper.AuthenticationData `swaggerignore:"true"`
}

type MultipleIdRequest struct {
	Ids      []int64                   `json:"ids"`
	AuthData helper.AuthenticationData `swaggerignore:"true"`
}

type CreateContentTypeRequest struct {
	Name string `json:"name"`

	IsCanView bool `json:"can_view"`

	IsCanCreate bool `json:"can_create"`

	IsCanEdit bool `json:"can_edit"`

	IsCanDelete bool `json:"can_delete"`
}

type CreateRoleRequest struct {
	Name string `json:"name"`

	PermissionIds []int64                   `json:"permission_ids"`
	AuthData      helper.AuthenticationData `swaggerignore:"true"`
}

type CreateUserRequest struct {
	Name     string                    `json:"name"`
	UserName string                    `json:"user_name"`
	Type     string                    `json:"type"`
	Password string                    `json:"password"`
	ClientId int64                     `json:"client_id"`
	RoleIds  []int64                   `json:"role_ids"`
	AuthData helper.AuthenticationData `swaggerignore:"true"`
}

type UpdateRoleRequest struct {
	Id            int64   `json:"id"`
	Name          string  `json:"name"`
	PermissionIds []int64 `json:"permission_ids"`
}

type CreateClientRequest struct {
	CompanyName string                    `json:"company_name"`
	AuthData    helper.AuthenticationData `swaggerignore:"true"`
}

type UpdateClientRequest struct {
	Id          int64                     `json:"id"`
	CompanyName string                    `json:"company_name"`
	AuthData    helper.AuthenticationData `swaggerignore:"true"`
}
