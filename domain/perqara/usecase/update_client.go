package usecase

import (
	"context"
	"errors"
	"log"
	"perqara/domain/constant"
	"perqara/domain/payload"
	"perqara/domain/request"
	"perqara/lib/helper"

	"google.golang.org/grpc/codes"
)

func (s *service) UpdateClient(ctx context.Context, req *request.UpdateClientRequest) error {
	var (
		clientUpdateRequest payload.UpdateClientPayload
	)
	isValidate, err := s.ValidateAccess(ctx, &helper.AuthenticationData{
		UserId:          req.AuthData.UserId,
		UserName:        req.AuthData.UserName,
		ClientId:        nil,
		ContentTypeName: constant.MenuContentTypeClient,
		Permissions:     constant.MenuPermissionUpdate,
	})
	if err != nil {
		return err
	}
	if !isValidate {
		return helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	if req == nil || req.Id == 0 {
		return helper.Error(codes.InvalidArgument, "Invalid Argument", errors.New("invalid argument"))
	}

	// Validate Request
	if len(req.CompanyName) > 0 {
		clientUpdateRequest.Name = req.CompanyName
	}

	clientUpdateRequest.ClientId = req.Id
	s.perqaraRepo.ForeignKeySwitch(0)
	defer s.perqaraRepo.ForeignKeySwitch(1)

	tx := s.externalLibUsecase.GormTxBegin()
	defer s.externalLibUsecase.GormTxRollback(tx)

	err = s.perqaraRepo.UpdateClient(ctx, tx, &clientUpdateRequest)
	if err != nil {
		log.Println(err)
		return err
	}

	s.externalLibUsecase.GormTxCommit(tx)

	return nil
}
