package usecase

import (
	"context"
	"errors"
	"log"
	"perqara/domain/payload"
	"perqara/domain/request"
	"perqara/lib/helper"
	"strconv"

	"google.golang.org/grpc/codes"
)

func (s *service) DeleteRole(ctx context.Context, req *request.SingleIdRequest) error {
	// Validate Required Request
	if req == nil || len(req.Id) == 0 {
		return helper.Error(codes.InvalidArgument, "Invalid Argument", errors.New("invalid argument"))
	}

	reqId, _ := strconv.ParseInt(req.Id, 10, 64)

	s.perqaraRepo.ForeignKeySwitch(0)
	defer s.perqaraRepo.ForeignKeySwitch(1)

	tx := s.externalLibUsecase.GormTxBegin()
	defer s.externalLibUsecase.GormTxRollback(tx)

	// Delete Roles
	err := s.perqaraRepo.DeleteRoles(ctx, tx, &payload.DeleteRolesPayload{
		RoleIds: []int64{reqId},
	})
	if err != nil {
		log.Println(err)
		return err
	}

	// Delete Role Permissions
	err = s.perqaraRepo.DeleteRolePermissions(ctx, tx, &payload.DeleteRolePermissionsPayload{
		RoleIds: []int64{reqId},
	})
	if err != nil {
		log.Println(err)
		return err
	}

	s.externalLibUsecase.GormTxCommit(tx)

	return nil
}
