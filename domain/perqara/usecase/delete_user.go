package usecase

import (
	"context"
	"errors"
	"log"
	"perqara/domain/constant"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"perqara/domain/request"
	"perqara/lib/helper"
	"strconv"

	"google.golang.org/grpc/codes"
)

func (s *service) DeleteUser(ctx context.Context, req *request.SingleIdRequest) error {
	// Validate Required Request
	if req == nil || len(req.Id) == 0 {
		return helper.Error(codes.InvalidArgument, "Invalid Argument", errors.New("invalid argument"))
	}
	isValidate, err := s.ValidateAccess(ctx, &helper.AuthenticationData{
		UserId:          req.AuthData.UserId,
		UserName:        req.AuthData.UserName,
		ClientId:        nil,
		ContentTypeName: constant.MenuContentTypeUser,
		Permissions:     constant.MenuPermissionDelete,
	})
	if err != nil {
		return err
	}
	if !isValidate {
		return helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	reqId, _ := strconv.ParseInt(req.Id, 10, 64)

	s.perqaraRepo.ForeignKeySwitch(0)
	defer s.perqaraRepo.ForeignKeySwitch(1)

	tx := s.externalLibUsecase.GormTxBegin()
	defer s.externalLibUsecase.GormTxRollback(tx)

	// Check is User Exist
	users, err := s.perqaraRepo.GetUsers(ctx, &payload.GetUsersPayload{
		UserIds: []int64{reqId},
	})
	if err != nil {
		log.Println(err)
		return err
	}
	if len(users) == 0 {
		return helper.Error(codes.Aborted, "User Not Found", errors.New("user not found"))
	}

	// Check if User External
	var (
		isExternal bool = false
		user       *model.User
	)
	user = users[0]
	if user.Type == constant.UserTypeExternalValue {
		isExternal = true
	}

	// Delete Users
	err = s.perqaraRepo.DeleteUsers(ctx, tx, &payload.DeleteUsersPayload{
		UserIds: []int64{reqId},
	})
	if err != nil {
		log.Println(err)
		return err
	}

	// Delete Role Permissions if User Internal
	if !isExternal {
		err = s.perqaraRepo.DeleteUserRoles(ctx, tx, &payload.DeleteUserRolesPayload{
			UserIds: []int64{reqId},
		})
		if err != nil {
			log.Println(err)
			return err
		}
	}

	s.externalLibUsecase.GormTxCommit(tx)

	return nil
}
