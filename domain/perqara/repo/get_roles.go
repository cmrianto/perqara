package repo

import (
	"context"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"perqara/lib/helper"
)

func (r *perqaraRepo) GetRoles(ctx context.Context, req *payload.GetRolesPayload) ([]*model.Role, error) {
	var result []*model.Role

	q := r.db.Model(&result).
		Where("deleted_at IS NULL").
		Order("created_at DESC")
	if req != nil {
		if len(req.RoleIds) > 0 {
			q.Where("id IN ?", req.RoleIds)
		}
		if req.Search != nil && len(*req.Search) > 0 {
			searchStr := "%" + helper.GetString(req.Search) + "%"
			q.Where("lower(name) like lower(?)", searchStr)
		}
		if req.Limit != nil {
			q.Limit(int(*req.Limit))
		}
		if req.Offset != nil {
			q.Offset(int(*req.Offset))
		}
		if req.LastCreatedAt != nil && len(*req.LastCreatedAt) > 0 {
			q.Where("created_at <= ?", helper.GetString(req.LastCreatedAt))
		}
	}
	err := q.Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}
