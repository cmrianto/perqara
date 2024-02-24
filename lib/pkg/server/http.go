package server

import (
	"fmt"
	"perqara/config"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Authorization, X-SKIP-AUTH, X-App-Name")
		c.Header("Access-Control-Allow-Methods", "DELETE, POST, HEAD, PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(CORSMiddleware())
	r.Use(gin.Recovery())
	r.Use(gin.ErrorLogger())

	return r
}

func Run(r *gin.Engine, cfg config.Config) error {
	// Start http server
	if err := r.Run(fmt.Sprintf(":%s", cfg.Application.ServerPort)); err != nil {
		return err
	}
	return nil
}
