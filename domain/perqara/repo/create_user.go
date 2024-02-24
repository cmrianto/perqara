package repo

import (
	"context"
	"perqara/domain/perqara/model"

	"gorm.io/gorm"
)

func (r *perqaraRepo) CreateUser(ctx context.Context, tx *gorm.DB, data *model.User) error {
	result := tx.Create(data)
	err := result.Error
	if err != nil {
		return err
	}

	result.Last(&data)
	return nil
}
