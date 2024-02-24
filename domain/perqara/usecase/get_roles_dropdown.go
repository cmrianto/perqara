package usecase

import (
	"context"
	"errors"
	"log"
	"perqara/domain/constant"
	"perqara/domain/payload"
	"perqara/domain/response"
	"perqara/lib/helper"
	"time"

	"google.golang.org/grpc/codes"
)

func (s *service) GetRolesDropdown(ctx context.Context, req *payload.GetRolesPayload) (*response.GetRolesResponse, error) {
	var (
		result response.GetRolesResponse
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
		tempRole := response.Role{
			Id:        r.Id,
			Name:      r.Name,
			CreatedAt: r.CreatedAt.Format(time.RFC3339),
		}
		result.Roles = append(result.Roles, tempRole)
	}
	return &result, nil
}
