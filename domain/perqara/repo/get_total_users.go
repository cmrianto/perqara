package repo

import (
	"context"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"perqara/lib/helper"
)

func (r *perqaraRepo) GetTotalUsers(ctx context.Context, req *payload.GetUsersPayload) (int64, error) {
	var result int64

	q := r.db.Model(&model.User{}).
		Where("deleted_at IS NULL")
	if req != nil {
		if req.Username != nil && len(*req.Username) > 0 {
			q.Where("user_name = ?", *req.Username)
		}
		if len(req.FilterUserTypes) > 0 {
			q.Where("type IN ?", req.FilterUserTypes)
		}
		if req.Search != nil && len(*req.Search) > 0 {
			searchStr := "%" + helper.GetString(req.Search) + "%"
			q.Where("lower(name) like lower(?)", searchStr)
		}
		if req.LastCreatedAt != nil && len(*req.LastCreatedAt) > 0 {
			q.Where("created_at <= ?", helper.GetString(req.LastCreatedAt))
		}
	}
	err := q.Count(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}
