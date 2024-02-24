package external_interface

import (
	"context"

	"gorm.io/gorm"
)

type ExternalLibInterface interface {
	// catch recover
	CatchRecover(ctx context.Context)

	// gorm
	GormTxBegin() *gorm.DB
	GormTxRollback(tx *gorm.DB)
	GormTxCommit(tx *gorm.DB)
}
