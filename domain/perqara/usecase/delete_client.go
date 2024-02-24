package usecase

import (
	"context"
	"errors"
	"log"
	"perqara/domain/constant"
	"perqara/domain/payload"
	"perqara/domain/request"
	"perqara/lib/helper"
	"strconv"

	"google.golang.org/grpc/codes"
)

func (s *service) DeleteClient(ctx context.Context, req *request.SingleIdRequest) error {
	// Validate Required Request
	if req == nil || len(req.Id) == 0 {
		return helper.Error(codes.InvalidArgument, "Invalid Argument", errors.New("invalid argument"))
	}
	isValidate, err := s.ValidateAccess(ctx, &helper.AuthenticationData{
		UserId:          req.AuthData.UserId,
		UserName:        req.AuthData.UserName,
		ClientId:        nil,
		ContentTypeName: constant.MenuContentTypeClient,
		Permissions:     constant.MenuPermissionDelete,
	})
	if err != nil {
		return err
	}
	if !isValidate {
		return helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	reqId, _ := strconv.ParseInt(req.Id, 10, 64)

	clients, err := s.perqaraRepo.GetClients(ctx, &payload.GetClientsPayload{
		ClientIds: &[]int64{reqId},
	})
	if err != nil {
		return err
	}
	if clients == nil || len(*clients) == 0 {
		return nil
	}

	s.perqaraRepo.ForeignKeySwitch(0)
	defer s.perqaraRepo.ForeignKeySwitch(1)

	tx := s.externalLibUsecase.GormTxBegin()
	defer s.externalLibUsecase.GormTxRollback(tx)

	// Delete  clients
	err = s.perqaraRepo.DeleteClients(ctx, tx, &payload.DeleteClientsPayload{
		ClientIds: []int64{reqId},
	})
	if err != nil {
		log.Println(err)
		return err
	}

	s.externalLibUsecase.GormTxCommit(tx)

	return nil
}
