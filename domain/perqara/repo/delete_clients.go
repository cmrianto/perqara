package repo

import (
	"context"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"time"

	"gorm.io/gorm"
)

func (r *perqaraRepo) DeleteClients(ctx context.Context, tx *gorm.DB, req *payload.DeleteClientsPayload) error {
	if req == nil ||
		len(req.ClientIds) == 0 {
		return nil
	}

	timeNow := time.Now().UTC()
	q := tx.Model(&model.Client{})
	if req != nil {
		if len(req.ClientIds) > 0 {
			q.Where("id IN ?", req.ClientIds)
		}
	}

	result := q.Updates(model.Client{
		UpdatedAt: &timeNow,
		DeletedAt: &timeNow,
	})
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}
