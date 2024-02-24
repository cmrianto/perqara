package repo

import (
	"context"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"perqara/lib/helper"
)

func (r *perqaraRepo) GetUsers(ctx context.Context, req *payload.GetUsersPayload) ([]*model.User, error) {
	var result []*model.User

	q := r.db.Model(&result).
		Where("deleted_at IS NULL").
		Order("created_at DESC")
	if req != nil {
		if len(req.UserIds) > 0 {
			q.Where("id IN ?", req.UserIds)
		}
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
