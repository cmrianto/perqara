package application

import (
	"context"
	"log"
	"perqara/config"
	"perqara/lib/pkg/database/mysql_gorm"
	"perqara/lib/pkg/redis"

	"github.com/urfave/cli"
)

func Setup(cfg *config.Config, c *cli.Context) (*Application, error) {
	app := new(Application)

	app.ServiceName = c.String("name")
	baseInit := []func(*Application) error{
		initDatabase(cfg),
		initRedis(cfg),
	}

	if err := runInit(baseInit...)(app); err != nil {
		return app, err
	}

	return app, nil
}

func runInit(appFuncs ...func(*Application) error) func(*Application) error {
	return func(app *Application) error {
		app.Context = context.Background()
		for _, fn := range appFuncs {
			if err := fn(app); err != nil {
				return err
			}
		}
		return nil
	}
}

func initDatabase(cfg *config.Config) func(*Application) error {
	return func(app *Application) error {
		db, err := mysql_gorm.NewMysqlORM(cfg)
		if err != nil {
			return err
		}

		app.DbClients = map[string]*DbClient{
			"read": {
				Type:       ReadConnection,
				SqlAdapter: db,
			},
			"write": {
				Type:       WriteConnection,
				SqlAdapter: db,
			},
		}
		log.Println("init database done")
		return nil
	}
}

func initRedis(cfg *config.Config) func(*Application) error {
	return func(app *Application) error {
		rdb, err := redis.NewRedisClient(cfg)
		if err != nil {
			return err
		}
		app.RedisClient = rdb
		return nil
	}
}
