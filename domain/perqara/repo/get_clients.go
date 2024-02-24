package repo

import (
	"context"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"perqara/lib/helper"
)

func (r *perqaraRepo) GetClients(ctx context.Context, req *payload.GetClientsPayload) (*[]model.Client, error) {
	var result []model.Client
	q := r.db.Model(&result).
		Where("clients.deleted_at IS NULL").
		Order("clients.created_at DESC")
	if req != nil {
		if req.ClientIds != nil && len(*req.ClientIds) > 0 {
			q.Where("clients.id IN ?", *req.ClientIds)
		}
		if req.Search != nil && len(*req.Search) > 0 {
			searchStr := "%" + helper.GetString(req.Search) + "%"
			q.Where("lower(clients.name) like lower(?)", searchStr)
		}
		if req.Limit != nil {
			q.Limit(int(*req.Limit))
		}
		if req.Offset != nil {
			q.Offset(int(*req.Offset))
		}
		if req.LastCreatedAt != nil && len(*req.LastCreatedAt) > 0 {
			q.Where("clients.created_at <= ?", helper.GetString(req.LastCreatedAt))
		}
	}
	err := q.Find(&result).Error
	if err != nil {
		return &result, err
	}

	return &result, nil
}
