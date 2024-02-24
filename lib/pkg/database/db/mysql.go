package db

import (
	"database/sql"
	"fmt"

	"perqara/config"
	"perqara/lib/pkg/logger"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

func NewMysqlDB(cfg *config.Config) (*sql.DB, error) {
	sqltrace.Register("mysql", &pq.Driver{})
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Database,
	)

	logger.GetLogger().Info(fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Database,
	))

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = PingDB(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func PingDB(db *sql.DB) error {
	_, err := db.Exec("SELECT 1")
	return err
}
