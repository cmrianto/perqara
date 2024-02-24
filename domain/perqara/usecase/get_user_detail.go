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

func (s *service) GetUserDetail(ctx context.Context, req *payload.GetSingleIdPayload) (*response.User, error) {
	var (
		result response.User

		user model.User

		roleIds []int64
		roles   []*model.Role

		rolePermissions []*model.RolePermission

		permissionIds []int64
		permissions   *[]model.Permission

		contentTypeIds []int64
		contentTypes   *[]model.ContentType
	)
	isValidate, err := s.ValidateAccess(ctx, &helper.AuthenticationData{
		UserId:          req.AuthData.UserId,
		UserName:        req.AuthData.UserName,
		ClientId:        nil,
		ContentTypeName: constant.MenuContentTypeUser,
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

	users, err := s.perqaraRepo.GetUsers(ctx, &payload.GetUsersPayload{
		UserIds: []int64{*req.Id},
	})
	if err != nil {
		log.Println(err)
		return &result, err
	}
	if len(users) == 0 {
		return &result, nil
	}

	for _, u := range users {
		if u != nil {
			user = *u
			break
		}
	}

	userRoles, err := s.perqaraRepo.GetUserRoles(ctx, &payload.GetUserRolesPayload{
		UserIds: &[]int64{user.Id},
	})
	if err != nil {
		log.Println(err)
		return &result, err
	}

	roleMapping := make(map[int64][]model.RolePermission)
	rolePermissionMapping := make(map[int64][]model.Permission)
	permissionMapping := make(map[int64][]model.ContentType)
	if len(userRoles) > 0 {
		for _, ur := range userRoles {
			if ur != nil {
				roleIds = append(roleIds, ur.RoleId)
			}
		}
		if len(roleIds) > 0 {
			roles, err = s.perqaraRepo.GetRoles(ctx, &payload.GetRolesPayload{
				RoleIds: roleIds,
			})
			if err != nil {
				log.Println(err)
				return &result, err
			}
			if len(roles) > 0 {
				rolePermissions, err = s.perqaraRepo.GetRolePermissions(ctx, &payload.GetRolePermissionsPayload{
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
					for _, r := range roles {
						for _, rp := range rolePermissions {
							if rp != nil && r.Id == rp.RoleId {
								roleMapping[r.Id] = append(roleMapping[r.Id], *rp)
							}
						}
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
						for _, rp := range rolePermissions {
							for _, p := range *permissions {
								tempPermission := p
								if rp.PermissionId == tempPermission.Id {
									rolePermissionMapping[rp.RoleId] = append(rolePermissionMapping[rp.RoleId], tempPermission)
								}
							}
						}
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

					for _, p := range *permissions {
						for _, ct := range *contentTypes {
							tempData := ct
							if p.ContentTypeId == tempData.Id {
								permissionMapping[p.Id] = append(permissionMapping[p.Id], ct)
							}
						}
					}
				}
			}
		}
	}
	result.Id = user.Id
	result.Name = user.Name
	result.UserName = user.UserName
	if user.ClientId.Valid && user.ClientId.Int64 > 0 {
		result.ClientId = user.ClientId.Int64
		result.Type = constant.UserTypeName[user.ClientId.Int64]
	} else {
		result.Type = constant.UserTypeInternal
	}

	if len(roles) > 0 {
		for _, r := range roles {
			if r != nil {
				tempRole := response.Role{
					Id:          r.Id,
					Name:        r.Name,
					CreatedAt:   r.CreatedAt.Format(time.RFC3339),
					Permissions: []response.Permission{},
				}
				currentRolePermissions, exist := roleMapping[r.Id]
				if exist {
					if len(currentRolePermissions) > 0 {
						for _, crp := range currentRolePermissions {
							currentPermission, exist := rolePermissionMapping[crp.RoleId]
							if exist {
								if len(currentPermission) > 0 {
									for i := range currentPermission {
										currentContentType, exist := permissionMapping[currentPermission[i].Id]
										if exist {
											if len(currentContentType) > 0 {
												for _, cct := range currentContentType {
													tempRole.Permissions = append(tempRole.Permissions, response.Permission{
														Id:              currentPermission[i].Id,
														ContentTypeId:   cct.Id,
														Name:            currentPermission[i].Name,
														ContentTypeName: cct.Name,
													})
												}
											}
										}
									}
								}
							}
						}
					}
				}
				result.Roles = append(result.Roles, tempRole)
			}
		}
	}

	return &result, nil
}
