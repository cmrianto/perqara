package usecase

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"perqara/domain/constant"
	"perqara/domain/perqara/model"
	"perqara/domain/request"
	"perqara/lib/helper"
	"time"

	"google.golang.org/grpc/codes"
)

func (s *service) CreateUser(ctx context.Context, req *request.CreateUserRequest) error {
	var (
		userData model.User

		userRoleDatas []*model.UserRole
	)
	isValidate, err := s.ValidateAccess(ctx, &helper.AuthenticationData{
		UserId:          req.AuthData.UserId,
		UserName:        req.AuthData.UserName,
		ClientId:        nil,
		ContentTypeName: constant.MenuContentTypeUser,
		Permissions:     constant.MenuPermissionCreate,
	})
	if err != nil {
		return err
	}
	if !isValidate {
		return helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	timeNow := time.Now().UTC()
	// Validate Required User Input Request
	if req == nil ||
		len(req.Name) == 0 ||
		len(req.UserName) == 0 ||
		len(req.Password) == 0 ||
		len(req.Type) == 0 {
		return helper.Error(codes.InvalidArgument, "Invalid Argument", errors.New("invalid argument"))
	}

	if helper.IsHaveSpace(req.UserName) || helper.IsHaveTab(req.UserName) {
		return helper.Error(codes.InvalidArgument, "Invalid Username", errors.New("invalid username"))
	}

	userTypeValue, exist := constant.UserTypeValue[req.Type]
	if !exist {
		return helper.Error(codes.InvalidArgument, "Invalid User Type", errors.New("invalid user type"))
	}

	if req.Type == constant.UserTypeExternal {
		if req.ClientId == 0 {
			return helper.Error(codes.InvalidArgument, "Invalid Client ID", errors.New("invalid client id"))
		}
	}

	pw, err := helper.HashPassword(req.Password)
	if err != nil {
		return err
	}

	var (
		isExternal bool = false
	)

	// Prepare Insert User
	userData.Name = req.Name
	userData.UserName = req.UserName
	userData.Password = pw
	userData.Type = userTypeValue
	if userTypeValue == constant.UserTypeExternalValue {
		userData.ClientId = sql.NullInt64{Int64: req.ClientId, Valid: true}
		isExternal = true
	}
	userData.CreatedAt = timeNow
	userData.UpdatedAt = &timeNow

	if !isExternal && len(req.RoleIds) == 0 {
		return helper.Error(codes.InvalidArgument, "Invalid Argument", errors.New("invalid argument"))
	}

	s.perqaraRepo.ForeignKeySwitch(0)
	defer s.perqaraRepo.ForeignKeySwitch(1)

	tx := s.externalLibUsecase.GormTxBegin()
	defer s.externalLibUsecase.GormTxRollback(tx)

	err = s.perqaraRepo.CreateUser(ctx, tx, &userData)
	if err != nil {
		log.Println(err)
		return err
	}

	// Insert User Roles for Internal User Type
	if !isExternal {
		for _, ur := range req.RoleIds {
			userRoleDatas = append(userRoleDatas, &model.UserRole{
				UserId:    userData.Id,
				RoleId:    ur,
				CreatedAt: timeNow,
				UpdatedAt: &timeNow,
			})
		}

		err = s.perqaraRepo.CreateUserRoles(ctx, tx, userRoleDatas)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	s.externalLibUsecase.GormTxCommit(tx)
	return nil
}
