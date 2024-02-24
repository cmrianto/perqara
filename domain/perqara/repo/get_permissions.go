package repo

import (
	"context"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"perqara/lib/helper"
)

func (r *perqaraRepo) GetPermissions(ctx context.Context, req *payload.GetPermissionsPayload) (*[]model.Permission, error) {
	var result []model.Permission

	q := r.db.Model(&result).
		Where("deleted_at IS NULL").
		Order("created_at DESC")
	if req != nil {
		if len(req.PermissionIds) > 0 {
			q.Where("id IN ?", req.PermissionIds)
		}
		if req.Search != nil && len(*req.Search) > 0 {
			searchStr := "%" + helper.GetString(req.Search) + "%"
			q.Where("lower(name) like lower(?)", searchStr)
		}
		if len(req.ContentTypeIds) > 0 {
			q.Where("content_type_id IN ?", req.ContentTypeIds)
		}
		if len(req.PermissionNames) > 0 {
			q.Where("name IN ?", req.PermissionNames)
		}
	}
	err := q.Find(&result).Error
	if err != nil {
		return &result, err
	}

	return &result, nil
}
