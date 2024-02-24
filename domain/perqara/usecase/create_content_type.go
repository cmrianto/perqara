package usecase

import (
	"context"
	"errors"
	"log"
	"perqara/domain/perqara/model"
	"perqara/domain/request"
	"perqara/lib/helper"
	"time"

	"google.golang.org/grpc/codes"
)

func (s *service) CreateContentType(ctx context.Context, req *request.CreateContentTypeRequest) error {
	var (
		contentTypeData model.ContentType
		permissionDatas []model.Permission
	)
	timeNow := time.Now().UTC()
	// Validate Required Permission Input Request
	if req == nil ||
		len(req.Name) == 0 {
		return helper.Error(codes.InvalidArgument, "Invalid Argument", errors.New("invalid argument"))
	}

	// Prepare Insert Permission
	contentTypeData.Name = req.Name
	contentTypeData.CreatedAt = timeNow
	contentTypeData.UpdatedAt = &timeNow

	s.perqaraRepo.ForeignKeySwitch(0)
	defer s.perqaraRepo.ForeignKeySwitch(1)

	tx := s.externalLibUsecase.GormTxBegin()
	defer s.externalLibUsecase.GormTxRollback(tx)

	// Insert Content Type
	err := s.perqaraRepo.CreateContentType(ctx, tx, &contentTypeData)
	if err != nil {
		log.Println(err)
		return err
	}

	// Prepare Insert Permissions
	if req.IsCanCreate {
		permissionDatas = append(permissionDatas, model.Permission{
			ContentTypeId: contentTypeData.Id,
			Name:          "can create",
			CreatedAt:     timeNow,
			UpdatedAt:     &timeNow,
		})
	}
	if req.IsCanDelete {
		permissionDatas = append(permissionDatas, model.Permission{
			ContentTypeId: contentTypeData.Id,
			Name:          "can delete",
			CreatedAt:     timeNow,
			UpdatedAt:     &timeNow,
		})
	}
	if req.IsCanEdit {
		permissionDatas = append(permissionDatas, model.Permission{
			ContentTypeId: contentTypeData.Id,
			Name:          "can edit",
			CreatedAt:     timeNow,
			UpdatedAt:     &timeNow,
		})
	}
	if req.IsCanView {
		permissionDatas = append(permissionDatas, model.Permission{
			ContentTypeId: contentTypeData.Id,
			Name:          "can view",
			CreatedAt:     timeNow,
			UpdatedAt:     &timeNow,
		})
	}

	// Create Permissions
	if len(permissionDatas) > 0 {
		err := s.perqaraRepo.CreatePermissions(ctx, tx, &permissionDatas)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	s.externalLibUsecase.GormTxCommit(tx)
	return nil
}
