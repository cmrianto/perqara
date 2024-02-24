package usecase

import (
	"perqara/config"
	external_interface "perqara/lib/external_lib"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type usecase struct {
	env         string
	db          *gorm.DB
	config      *config.Config
	redisClient redis.UniversalClient
}

type Dependencies struct {
	Env         string
	Db          *gorm.DB
	Config      *config.Config
	RedisClient redis.UniversalClient
}

func NewUsecase(deps Dependencies) external_interface.ExternalLibInterface {
	return &usecase{
		env:         deps.Env,
		db:          deps.Db,
		config:      deps.Config,
		redisClient: deps.RedisClient,
	}
}
