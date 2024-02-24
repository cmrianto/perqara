package repo

import (
	"context"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"time"

	"gorm.io/gorm"
)

func (r *perqaraRepo) UpdateClient(ctx context.Context, tx *gorm.DB, req *payload.UpdateClientPayload) error {
	if req == nil || req.ClientId == 0 {
		return nil
	}

	var updateRequest model.Client

	q := tx.Model(&model.Client{}).
		Where("id = ?", req.ClientId)

	timeNow := time.Now().UTC()
	updateRequest.UpdatedAt = &timeNow
	if len(req.Name) > 0 {
		updateRequest.Name = req.Name
	}

	result := q.Updates(updateRequest)
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}
