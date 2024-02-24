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

func (s *service) GetRoleDetail(ctx context.Context, req *payload.GetSingleIdPayload) (*response.Role, error) {
	var (
		result response.Role

		role model.Role

		permissionIds []int64

		contentTypeIds []int64
		contentTypes   *[]model.ContentType
	)
	isValidate, err := s.ValidateAccess(ctx, &helper.AuthenticationData{
		UserId:          req.AuthData.UserId,
		UserName:        req.AuthData.UserName,
		ClientId:        nil,
		ContentTypeName: constant.MenuContentTypeAccess,
		Permissions:     constant.MenuPermissionDetail,
	})
	if err != nil {
		return &result, err
	}
	if !isValidate {
		return &result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	if req.Id == nil {
		return &result, helper.Error(codes.InvalidArgument, "Invalid Argument", errors.New("invalid argument"))
	}

	roles, err := s.perqaraRepo.GetRoles(ctx, &payload.GetRolesPayload{
		RoleIds: []int64{*req.Id},
	})
	if err != nil {
		log.Println(err)
		return &result, err
	}
	if len(roles) == 0 {
		return &result, nil
	}
	for _, r := range roles {
		role = *r
		break
	}

	rolePermissions, err := s.perqaraRepo.GetRolePermissions(ctx, &payload.GetRolePermissionsPayload{
		RoleIds: []int64{role.Id},
	})
	if err != nil {
		log.Println(err)
		return &result, err
	}
	if len(rolePermissions) == 0 {
		return &result, nil
	}

	if len(rolePermissions) > 0 {
		for _, rp := range rolePermissions {
			permissionIds = append(permissionIds, rp.PermissionId)
		}
	}

	permissions, err := s.perqaraRepo.GetPermissions(ctx, &payload.GetPermissionsPayload{
		PermissionIds: permissionIds,
	})
	if err != nil {
		log.Println(err)
		return &result, err
	}
	if permissions == nil || len(*permissions) == 0 {
		return &result, nil
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

	tempPermissions := []response.Permission{}
	for _, rp := range rolePermissions {
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
	result.Id = role.Id
	result.Name = role.Name
	result.CreatedAt = role.CreatedAt.Format(time.RFC3339)
	result.Permissions = tempPermissions

	return &result, nil
}
