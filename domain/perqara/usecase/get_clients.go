package usecase

import (
	"context"
	"errors"
	"perqara/domain/constant"
	"perqara/domain/payload"
	"perqara/domain/response"
	"perqara/lib/helper"
	"time"

	"google.golang.org/grpc/codes"
)

func (s *service) GetClients(ctx context.Context, req *payload.GetClientsPayload) (*response.GetClientsResponse, error) {
	var (
		result response.GetClientsResponse

		clientIds []int64
	)
	isValidate, err := s.ValidateAccess(ctx, &helper.AuthenticationData{
		UserId:          req.AuthData.UserId,
		UserName:        req.AuthData.UserName,
		ClientId:        nil,
		ContentTypeName: constant.MenuContentTypeClient,
		Permissions:     constant.MenuPermissionList,
	})
	if err != nil {
		return &result, err
	}
	if !isValidate {
		return &result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	clients, err := s.perqaraRepo.GetClients(ctx, req)
	if err != nil {
		return &result, err
	}
	if clients == nil || len(*clients) == 0 {
		return &result, nil
	}
	for _, c := range *clients {
		clientIds = append(clientIds, c.Id)
	}

	for _, c := range *clients {
		tempResult := response.Client{
			Id:          c.Id,
			CompanyName: c.Name,
			CreatedAt:   c.CreatedAt.Format(time.RFC3339),
		}
		result.Clients = append(result.Clients, tempResult)
	}

	return &result, nil
}
