package application

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Application struct {
	Context     context.Context
	ServiceName string

	DbClients   map[string]*DbClient
	RedisClient redis.UniversalClient

	HttpServer *gin.Engine
}

type DbConnectionType uint32

const (
	ReadConnection DbConnectionType = iota
	WriteConnection
)

type DbClient struct {
	Type       DbConnectionType
	SqlAdapter *gorm.DB
}
