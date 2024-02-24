package usecase

import (
	"context"
	"errors"
	"log"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"perqara/domain/request"
	"perqara/lib/helper"
	"time"

	"google.golang.org/grpc/codes"
)

func (s *service) UpdateRole(ctx context.Context, req *request.UpdateRoleRequest) error {
	var (
		roleUpdateRequest     payload.UpdateRolePayload
		isClearRolePermission bool = false

		rolePermissionDatas []model.RolePermission
	)
	if req == nil || req.Id == 0 || len(req.Name) == 0 {
		return helper.Error(codes.InvalidArgument, "Invalid Argument", errors.New("invalid argument"))
	}
	timeNow := time.Now().UTC()

	roleUpdateRequest.Id = &req.Id
	roleUpdateRequest.Name = &req.Name

	if len(req.PermissionIds) > 0 {
		isClearRolePermission = true
	}

	s.perqaraRepo.ForeignKeySwitch(0)
	defer s.perqaraRepo.ForeignKeySwitch(1)

	tx := s.externalLibUsecase.GormTxBegin()
	defer s.externalLibUsecase.GormTxRollback(tx)

	err := s.perqaraRepo.UpdateRole(ctx, tx, &roleUpdateRequest)
	if err != nil {
		log.Println(err)
		return err
	}

	if isClearRolePermission {
		err = s.perqaraRepo.DeleteRolePermissions(ctx, tx, &payload.DeleteRolePermissionsPayload{
			RoleIds: []int64{req.Id},
		})
		if err != nil {
			log.Println(err)
			return err
		}

		for _, pi := range req.PermissionIds {
			rolePermissionDatas = append(rolePermissionDatas, model.RolePermission{
				RoleId:       req.Id,
				PermissionId: pi,
				CreatedAt:    timeNow,
				UpdatedAt:    &timeNow,
			})
		}
		err := s.perqaraRepo.CreateRolePermissions(ctx, tx, &rolePermissionDatas)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	s.externalLibUsecase.GormTxCommit(tx)

	return nil
}
