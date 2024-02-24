package usecase

import "gorm.io/gorm"

func (e *usecase) GormTxBegin() *gorm.DB {
	return e.db.Begin()
}

func (e *usecase) GormTxRollback(tx *gorm.DB) {
	tx.Rollback()
}

func (e *usecase) GormTxCommit(tx *gorm.DB) {
	tx.Commit()
}
