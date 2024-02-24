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

func (s *service) GetTotalRoles(ctx context.Context, req *payload.GetRolesPayload) (*response.GetTotalResponse, error) {
	var (
		result response.GetTotalResponse
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

	totalRoles, err := s.perqaraRepo.GetTotalRoles(ctx, req)
	if err != nil {
		return &result, err
	}

	result.Total = totalRoles

	return &result, nil
}
