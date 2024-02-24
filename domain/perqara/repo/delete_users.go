package repo

import (
	"context"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"time"

	"gorm.io/gorm"
)

func (r *perqaraRepo) DeleteUsers(ctx context.Context, tx *gorm.DB, req *payload.DeleteUsersPayload) error {
	if req == nil || len(req.UserIds) == 0 {
		return nil
	}

	timeNow := time.Now().UTC()
	q := tx.Model(&model.User{})
	if req != nil {
		if req.UserIds != nil && len(req.UserIds) > 0 {
			q.Where("id IN ?", req.UserIds)
		}
	}

	result := q.Updates(model.User{
		UpdatedAt: &timeNow,
		DeletedAt: &timeNow,
	})
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}
