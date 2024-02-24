package usecase

import (
	"context"
	"log"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"perqara/domain/response"
	"time"
)

func (s *service) GetContentTypes(ctx context.Context, req *payload.GetContentTypesPayload) (*response.GetContentTypesResponse, error) {
	var (
		result response.GetContentTypesResponse

		ctIds       []int64
		permissions *[]model.Permission
	)
	contentTypes, err := s.perqaraRepo.GetContentTypes(ctx, req)
	if err != nil {
		return &result, err
	}

	if contentTypes == nil || len(*contentTypes) == 0 {
		return &result, nil
	}

	for _, ct := range *contentTypes {
		ctIds = append(ctIds, ct.Id)
	}

	if len(ctIds) > 0 {
		permissions, err = s.perqaraRepo.GetPermissions(ctx, &payload.GetPermissionsPayload{
			ContentTypeIds: ctIds,
		})
		if err != nil {
			log.Println(err)
			return &result, err
		}
	}

	for _, ct := range *contentTypes {
		tempContentType := response.ContentType{
			Id:        ct.Id,
			Name:      ct.Name,
			CreatedAt: ct.CreatedAt.Format(time.RFC3339),
		}
		if permissions != nil && len(*permissions) > 0 {
			tempPermissions := []response.Permission{}
			for _, p := range *permissions {
				if ct.Id == p.ContentTypeId {
					tempPermissions = append(tempPermissions, response.Permission{
						Id:              p.Id,
						Name:            p.Name,
						ContentTypeId:   p.ContentTypeId,
						ContentTypeName: ct.Name,
					})
				}
			}
			tempContentType.Permissions = tempPermissions
		}
		result.ContentTypes = append(result.ContentTypes, tempContentType)
	}

	return &result, nil
}
