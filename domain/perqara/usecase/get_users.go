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

	"google.golang.org/grpc/codes"
)

func (s *service) GetUsers(ctx context.Context, req *payload.GetUsersPayload) (*response.GetUsersResponse, error) {
	var (
		result response.GetUsersResponse

		userIds []int64
		roleIds []int64

		roles []*model.Role
	)
	isValidate, err := s.ValidateAccess(ctx, &helper.AuthenticationData{
		UserId:          req.AuthData.UserId,
		UserName:        req.AuthData.UserName,
		ClientId:        nil,
		ContentTypeName: constant.MenuContentTypeUser,
		Permissions:     constant.MenuPermissionList,
	})
	if err != nil {
		return &result, err
	}
	if !isValidate {
		return &result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	users, err := s.perqaraRepo.GetUsers(ctx, req)
	if err != nil {
		log.Println(err)
		return &result, err
	}

	if len(users) == 0 {
		return &result, nil
	}

	for _, u := range users {
		if u != nil {
			userIds = append(userIds, u.Id)
		}
	}

	userRoleMap := make(map[int64][]*model.UserRole)
	userRoles, err := s.perqaraRepo.GetUserRoles(ctx, &payload.GetUserRolesPayload{
		UserIds: &userIds,
	})
	if err != nil {
		log.Println(err)
		return &result, err
	}

	if len(userRoles) > 0 {
		var urModel *model.UserRole
		userRoleMap = urModel.MakeMapArrayByUserId(userRoles)
		for _, ur := range userRoles {
			if ur != nil {
				roleIds = append(roleIds, ur.RoleId)
			}
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

	}

	for _, u := range users {
		roleResponse := []response.Role{}
		_, urExist := userRoleMap[u.Id]
		if urExist {
			for _, ur := range userRoleMap[u.Id] {
				for _, r := range roles {
					if ur.RoleId == r.Id {
						tempRole := response.Role{
							Id:          r.Id,
							Name:        r.Name,
							Permissions: []response.Permission{},
						}
						roleResponse = append(roleResponse, tempRole)
					}
				}
			}
		}
		result.Users = append(result.Users, response.User{
			Id:       u.Id,
			Name:     u.Name,
			UserName: u.UserName,
			Type:     constant.UserTypeName[u.Type],
			ClientId: u.ClientId.Int64,
			Roles:    roleResponse,
		})
	}
	return &result, nil
}
