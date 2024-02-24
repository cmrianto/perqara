package repo

import (
	"context"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"time"

	"gorm.io/gorm"
)

func (r *perqaraRepo) UpdateRole(ctx context.Context, tx *gorm.DB, req *payload.UpdateRolePayload) error {
	if req == nil || req.Id == nil || *req.Id == 0 {
		return nil
	}

	var updateRequest model.Role

	q := tx.Model(&model.Role{}).
		Where("id = ?", *req.Id)

	timeNow := time.Now().UTC()
	updateRequest.UpdatedAt = &timeNow
	if req.Name != nil && len(*req.Name) > 0 {
		updateRequest.Name = *req.Name
	}

	result := q.Updates(updateRequest)
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}
