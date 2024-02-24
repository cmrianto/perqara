package mysql_gorm

import (
	"database/sql"
	"fmt"

	"perqara/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlORM(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.Database.User, cfg.Database.Password,
		cfg.Database.Host, cfg.Database.Port,
		cfg.Database.Database,
	)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	return db, err
}
