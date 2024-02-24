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

func (s *service) CreateRole(ctx context.Context, req *request.CreateRoleRequest) error {
	var (
		roleData            model.Role
		rolePermissionDatas []model.RolePermission
	)
	isValidate, err := s.ValidateAccess(ctx, &helper.AuthenticationData{
		UserId:          req.AuthData.UserId,
		UserName:        req.AuthData.UserName,
		ClientId:        nil,
		ContentTypeName: constant.MenuContentTypeAccess,
		Permissions:     constant.MenuPermissionCreate,
	})
	if err != nil {
		return err
	}
	if !isValidate {
		return helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	timeNow := time.Now().UTC()
	// Validate Required Role Input Request
	if req == nil ||
		len(req.Name) == 0 {
		return helper.Error(codes.InvalidArgument, "Invalid Argument", errors.New("invalid argument"))
	}

	// Prepare Insert Role
	roleData.Name = req.Name
	roleData.CreatedAt = timeNow
	roleData.UpdatedAt = &timeNow

	s.perqaraRepo.ForeignKeySwitch(0)
	defer s.perqaraRepo.ForeignKeySwitch(1)

	tx := s.externalLibUsecase.GormTxBegin()
	defer s.externalLibUsecase.GormTxRollback(tx)

	err = s.perqaraRepo.CreateRole(ctx, tx, &roleData)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(req.PermissionIds) > 0 {
		for _, pi := range req.PermissionIds {
			rolePermissionDatas = append(rolePermissionDatas, model.RolePermission{
				RoleId:       roleData.Id,
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
