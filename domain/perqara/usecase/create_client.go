package usecase

import (
	"context"
	"errors"
	"log"
	"perqara/domain/constant"
	"perqara/domain/perqara/model"
	"perqara/domain/request"
	"perqara/lib/helper"
	"time"

	"google.golang.org/grpc/codes"
)

func (s *service) CreateClient(ctx context.Context, req *request.CreateClientRequest) error {
	var (
		clientData model.Client
	)
	isValidate, err := s.ValidateAccess(ctx, &helper.AuthenticationData{
		UserId:          req.AuthData.UserId,
		UserName:        req.AuthData.UserName,
		ClientId:        nil,
		ContentTypeName: constant.MenuContentTypeClient,
		Permissions:     constant.MenuPermissionCreate,
	})
	if err != nil {
		return err
	}
	if !isValidate {
		return helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	timeNow := time.Now().UTC()
	// Validate Required Client Input Request
	if req == nil ||
		len(req.CompanyName) == 0 {
		return helper.Error(codes.InvalidArgument, "Invalid Argument", errors.New("invalid argument"))
	}

	clientData.Name = req.CompanyName
	clientData.CreatedAt = timeNow
	clientData.UpdatedAt = &timeNow

	s.perqaraRepo.ForeignKeySwitch(0)
	defer s.perqaraRepo.ForeignKeySwitch(1)

	tx := s.externalLibUsecase.GormTxBegin()
	defer s.externalLibUsecase.GormTxRollback(tx)

	err = s.perqaraRepo.CreateClient(ctx, tx, &clientData)
	if err != nil {
		log.Println(err)
		return err
	}

	s.externalLibUsecase.GormTxCommit(tx)
	return nil
}
