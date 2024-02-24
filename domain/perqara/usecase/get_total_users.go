package usecase

import (
	"context"
	"errors"
	"perqara/domain/constant"
	"perqara/domain/payload"
	"perqara/domain/response"
	"perqara/lib/helper"

	"google.golang.org/grpc/codes"
)

func (s *service) GetTotalUsers(ctx context.Context, req *payload.GetUsersPayload) (*response.GetTotalResponse, error) {
	var (
		result response.GetTotalResponse
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

	totalUsers, err := s.perqaraRepo.GetTotalUsers(ctx, req)
	if err != nil {
		return &result, err
	}

	result.Total = totalUsers

	return &result, nil
}
