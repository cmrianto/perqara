package usecase

import (
	"context"
	"errors"
	"log"
	"perqara/domain/constant"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"perqara/domain/response"
	"perqara/lib/helper"
	"time"

	"google.golang.org/grpc/codes"
)

func (s *service) GetRoles(ctx context.Context, req *payload.GetRolesPayload) (*response.GetRolesResponse, error) {
	var (
		result response.GetRolesResponse

		roleIds        []int64
		permissionIds  []int64
		contentTypeIds []int64

		permissions  *[]model.Permission
		contentTypes *[]model.ContentType
	)
	isValidate, err := s.ValidateAccess(ctx, &helper.AuthenticationData{
		UserId:          req.AuthData.UserId,
		UserName:        req.AuthData.UserName,
		ClientId:        nil,
		ContentTypeName: constant.MenuContentTypeAccess,
		Permissions:     constant.MenuPermissionList,
	})
	if err != nil {
		return &result, err
	}
	if !isValidate {
		return &result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	roles, err := s.perqaraRepo.GetRoles(ctx, req)
	if err != nil {
		log.Println(err)
		return &result, err
	}

	if len(roles) == 0 {
		return &result, nil
	}

	for _, r := range roles {
		roleIds = append(roleIds, r.Id)
	}

	rolePermissions, err := s.perqaraRepo.GetRolePermissions(ctx, &payload.GetRolePermissionsPayload{
		RoleIds: roleIds,
	})
	if err != nil {
		log.Println(err)
		return &result, err
	}

	if len(rolePermissions) > 0 {
		for _, rp := range rolePermissions {
			permissionIds = append(permissionIds, rp.PermissionId)
		}
	}

	if len(permissionIds) > 0 {
		permissions, err = s.perqaraRepo.GetPermissions(ctx, &payload.GetPermissionsPayload{
			PermissionIds: permissionIds,
		})
		if err != nil {
			log.Println(err)
			return &result, err
		}

		if permissions != nil && len(*permissions) > 0 {
			for _, p := range *permissions {
				contentTypeIds = append(contentTypeIds, p.ContentTypeId)
			}
			if len(contentTypeIds) > 0 {
				contentTypes, err = s.perqaraRepo.GetContentTypes(ctx, &payload.GetContentTypesPayload{
					ContentTypeIds: contentTypeIds,
				})
				if err != nil {
					log.Println(err)
					return &result, err
				}
			}
		}
	}

	for _, r := range roles {
		tempRole := response.Role{
			Id:        r.Id,
			Name:      r.Name,
			CreatedAt: r.CreatedAt.Format(time.RFC3339),
		}
		if contentTypes != nil && len(*contentTypes) > 0 {
			tempPermissions := []response.Permission{}
			for _, rp := range rolePermissions {
				if rp.RoleId == r.Id {
					for _, p := range *permissions {
						if p.Id == rp.PermissionId {
							for _, ct := range *contentTypes {
								if ct.Id == p.ContentTypeId {
									tempPermissions = append(tempPermissions, response.Permission{
										Id:              p.Id,
										ContentTypeId:   ct.Id,
										Name:            p.Name,
										ContentTypeName: ct.Name,
									})
								}
							}
						}
					}
				}
			}
			tempRole.Permissions = tempPermissions
		}
		result.Roles = append(result.Roles, tempRole)
	}
	return &result, nil
}
