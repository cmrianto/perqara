package repo

import (
	"perqara/domain/perqara"

	"gorm.io/gorm"
)

type perqaraRepo struct {
	db *gorm.DB
}

func NewMysqlRepo(
	db *gorm.DB,
) perqara.PerqaraRepoInterface {
	return &perqaraRepo{
		db: db,
	}
}
