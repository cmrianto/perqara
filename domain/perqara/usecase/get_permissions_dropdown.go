package usecase

import (
	"context"
	"errors"
	"log"
	"perqara/domain/constant"
	"perqara/domain/payload"
	"perqara/domain/response"
	"perqara/lib/helper"

	"google.golang.org/grpc/codes"
)

func (s *service) GetPermissionsDropdown(ctx context.Context, req *payload.GetPermissionsPayload) (*response.GetPermissionsResponse, error) {
	var (
		result         response.GetPermissionsResponse
		contentTypeIds []int64
	)
	isValidate, err := s.ValidateAccess(ctx, &helper.AuthenticationData{
		UserId:          req.AuthData.UserId,
		UserName:        req.AuthData.UserName,
		ClientId:        nil,
		ContentTypeName: constant.MenuContentTypeAccess,
		Permissions:     constant.MenuPermissionCreate,
	})
	if err != nil {
		return &result, err
	}
	if !isValidate {
		return &result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	permissions, err := s.perqaraRepo.GetPermissions(ctx, req)
	if err != nil {
		log.Println(err)
		return &result, err
	}

	if permissions == nil || len(*permissions) == 0 {
		return &result, nil
	}

	for _, p := range *permissions {
		contentTypeIds = append(contentTypeIds, p.ContentTypeId)
	}

	contentTypes, err := s.perqaraRepo.GetContentTypes(ctx, &payload.GetContentTypesPayload{
		ContentTypeIds: contentTypeIds,
	})
	if err != nil {
		log.Println(err)
		return &result, err
	}

	for _, p := range *permissions {
		var contentTypeName string = ""
		if contentTypes != nil && len(*contentTypes) > 0 {
			for _, ct := range *contentTypes {
				if ct.Id == p.ContentTypeId {
					contentTypeName = ct.Name
				}
			}
		}
		result.Permissions = append(result.Permissions, response.Permission{
			Id:              p.Id,
			Name:            p.Name,
			ContentTypeId:   p.ContentTypeId,
			ContentTypeName: contentTypeName,
		})
	}
	return &result, nil
}
