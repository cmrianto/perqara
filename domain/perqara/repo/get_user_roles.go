package repo

import (
	"context"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
)

func (r *perqaraRepo) GetUserRoles(ctx context.Context, req *payload.GetUserRolesPayload) ([]*model.UserRole, error) {
	var result []*model.UserRole

	q := r.db.Model(&result).
		Where("deleted_at IS NULL").
		Order("created_at DESC")
	if req != nil {
		if len(*req.UserIds) > 0 {
			q.Where("user_id IN ?", *req.UserIds)
		}
	}

	err := q.Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}
