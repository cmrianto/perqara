package usecase

import (
	"context"
	"errors"
	"log"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"perqara/lib/helper"

	"google.golang.org/grpc/codes"
)

func (s *service) ValidateAccess(ctx context.Context, req *helper.AuthenticationData) (bool, error) {
	var (
		result bool = false

		user           *model.User
		roleIds        []int64
		permissionIds  []int64
		contentTypeIds []int64
	)
	if req == nil || req.UserId == 0 || len(req.Permissions) == 0 {
		return result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}
	users, err := s.perqaraRepo.GetUsers(ctx, &payload.GetUsersPayload{
		UserIds: []int64{req.UserId},
	})
	if err != nil {
		log.Println(err)
		return result, err
	}

	if len(users) == 0 {
		return result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}
	user = users[0]

	if req.ClientId != nil && *req.ClientId > 0 {
		clients, err := s.perqaraRepo.GetClients(ctx, &payload.GetClientsPayload{
			ClientIds: &[]int64{*req.ClientId},
		})
		if err != nil {
			return result, err
		}
		if clients == nil || len(*clients) == 0 {
			return result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
		}
		return true, nil
	}

	userRoles, err := s.perqaraRepo.GetUserRoles(ctx, &payload.GetUserRolesPayload{
		UserIds: &[]int64{user.Id},
	})
	if err != nil {
		log.Println(err)
		return result, err
	}
	if len(userRoles) == 0 {
		return result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	for _, ur := range userRoles {
		if ur != nil {
			tempUr := *ur
			roleIds = append(roleIds, tempUr.RoleId)
		}
	}
	if len(roleIds) == 0 {
		return result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	roles, err := s.perqaraRepo.GetRoles(ctx, &payload.GetRolesPayload{
		RoleIds: roleIds,
	})
	if err != nil {
		log.Println(err)
		return result, err
	}
	if len(roles) == 0 {
		return result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	roleIds = []int64{}
	for _, r := range roles {
		if r != nil {
			tempRole := *r
			roleIds = append(roleIds, tempRole.Id)
		}
	}
	if len(roleIds) == 0 {
		return result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	rolePermissions, err := s.perqaraRepo.GetRolePermissions(ctx, &payload.GetRolePermissionsPayload{
		RoleIds: roleIds,
	})
	if err != nil {
		log.Println(err)
		return result, err
	}
	if len(rolePermissions) == 0 {
		return result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	for _, rp := range rolePermissions {
		if rp != nil {
			tempRp := *rp
			permissionIds = append(permissionIds, tempRp.PermissionId)
		}
	}
	if len(permissionIds) == 0 {
		return result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	permissions, err := s.perqaraRepo.GetPermissions(ctx, &payload.GetPermissionsPayload{
		PermissionIds:   permissionIds,
		PermissionNames: req.Permissions,
	})
	if err != nil {
		log.Println(err)
		return result, err
	}
	if permissions == nil || len(*permissions) == 0 {
		return result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	for _, p := range *permissions {
		contentTypeIds = append(contentTypeIds, p.ContentTypeId)
	}
	if len(contentTypeIds) == 0 {
		return result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	contentTypes, err := s.perqaraRepo.GetContentTypes(ctx, &payload.GetContentTypesPayload{
		ContentTypeIds: contentTypeIds,
	})
	if err != nil {
		log.Println(err)
		return result, err
	}
	if contentTypes == nil || len(*contentTypes) == 0 {
		return result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}
	for _, ct := range *contentTypes {
		if ct.Name == req.ContentTypeName {
			result = true
			break
		}
	}

	return result, nil
}
