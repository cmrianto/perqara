package repo

import (
	"context"
	"perqara/domain/perqara/model"

	"gorm.io/gorm"
)

func (r *perqaraRepo) CreateUserRoles(ctx context.Context, tx *gorm.DB, data []*model.UserRole) error {
	result := tx.Create(&data)
	err := result.Error
	if err != nil {
		return err
	}

	result.Last(&data)
	return nil
}
