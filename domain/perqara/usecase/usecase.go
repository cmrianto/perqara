package usecase

import (
	"perqara/config"

	perqara "perqara/domain/perqara"

	external_interface "perqara/lib/external_lib"

	"github.com/go-redis/redis/v8"
)

type service struct {
	env                string
	perqaraRepo        perqara.PerqaraRepoInterface
	redisClient        redis.UniversalClient
	config             *config.Config
	externalLibUsecase external_interface.ExternalLibInterface
}

type Dependencies struct {
	Env                string
	PerqaraRepo        perqara.PerqaraRepoInterface
	RedisClient        redis.UniversalClient
	Config             *config.Config
	ExternalLibUsecase external_interface.ExternalLibInterface
}

func NewService(deps Dependencies) perqara.PerqaraUsecaseInterface {
	return &service{
		env:                deps.Env,
		perqaraRepo:        deps.PerqaraRepo,
		redisClient:        deps.RedisClient,
		config:             deps.Config,
		externalLibUsecase: deps.ExternalLibUsecase,
	}
}
