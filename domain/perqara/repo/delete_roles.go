package repo

import (
	"context"
	"perqara/domain/payload"
	"perqara/domain/perqara/model"
	"time"

	"gorm.io/gorm"
)

func (r *perqaraRepo) DeleteRoles(ctx context.Context, tx *gorm.DB, req *payload.DeleteRolesPayload) error {
	if req == nil || len(req.RoleIds) == 0 {
		return nil
	}

	timeNow := time.Now().UTC()
	q := tx.Model(&model.Role{})
	if req != nil {
		if req.RoleIds != nil && len(req.RoleIds) > 0 {
			q.Where("id IN ?", req.RoleIds)
		}
	}

	result := q.Updates(model.Role{
		UpdatedAt: &timeNow,
		DeletedAt: &timeNow,
	})
	err := result.Error
	if err != nil {
		return err
	}

	return nil
}
