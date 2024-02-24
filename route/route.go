package route

import (
	"context"
	"net/http"
	"perqara/application"
	"perqara/config"
	"perqara/docs"
	"perqara/domain/handler/http_handler"
	"perqara/domain/perqara/repo"
	"perqara/domain/perqara/usecase"
	"perqara/lib/helper"
	"perqara/lib/helper/http_response"
	"perqara/lib/pkg/database/mysql_gorm"

	"os"

	external_usecase "perqara/lib/external_lib/usecase"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type healthCheckResponse struct {
	Data string `json:"data"`
}

func JwtAuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.TokenValid(c, helper.TokenPayload{
			Secret: cfg.Auth.ApiSecret,
		})
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func SetupRouter(p *gin.Engine, cfg *config.Config, app *application.Application) {
	docs.SwaggerInfo.BasePath = ""
	p.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	ctx := context.Background()

	healthRes := healthCheckResponse{
		Data: "pong",
	}
	p.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, healthRes)
	})
	p.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, healthRes)
	})

	// Repo
	db, err := mysql_gorm.NewMysqlORM(cfg)
	if err != nil {
		panic(err)
	}
	mysqlRepo := repo.NewMysqlRepo(
		db,
	)

	// ExternalLibInterface
	externalLibInterface := external_usecase.NewUsecase(external_usecase.Dependencies{
		Env:         cfg.Application.Env,
		Db:          db,
		Config:      cfg,
		RedisClient: app.RedisClient,
	})

	// Usecase
	perqaraUsecase := usecase.NewService(usecase.Dependencies{
		Env:                cfg.Application.Env,
		PerqaraRepo:        mysqlRepo,
		RedisClient:        app.RedisClient,
		Config:             cfg,
		ExternalLibUsecase: externalLibInterface,
	})

	// Handler
	perqaraHandler := http_handler.NewPerqaraHandler(ctx, perqaraUsecase, *cfg)

	// Login
	p.POST("/login", ModeWrapper(perqaraHandler.Login))

	r := p.Group("")
	r.Use(JwtAuthMiddleware(cfg))

	// ContentType
	r.GET("/content-type", ModeWrapper(perqaraHandler.GetContentTypes))
	r.POST("/content-type", ModeWrapper(perqaraHandler.CreateContentType))

	// Permission
	r.GET("/permission/dropdown", ModeWrapper(perqaraHandler.GetPermissionsDropdown))

	// Role
	r.GET("/role", ModeWrapper(perqaraHandler.GetRoles))
	r.GET("/role/:id", ModeWrapper(perqaraHandler.GetRoleDetail))
	r.GET("/role/total", ModeWrapper(perqaraHandler.GetTotalRoles))
	r.GET("/role/dropdown", ModeWrapper(perqaraHandler.GetRolesDropdown))
	r.POST("/role", ModeWrapper(perqaraHandler.CreateRole))
	r.PUT("/role", ModeWrapper(perqaraHandler.UpdateRole))
	r.DELETE("/role/:id", ModeWrapper(perqaraHandler.DeleteRole))

	// User
	r.GET("/user", ModeWrapper(perqaraHandler.GetUsers))
	r.GET("/user/:id", ModeWrapper(perqaraHandler.GetUserDetail))
	r.GET("/user/total", ModeWrapper(perqaraHandler.GetTotalUsers))
	r.POST("/user", ModeWrapper(perqaraHandler.CreateUser))
	r.DELETE("/user/:id", ModeWrapper(perqaraHandler.DeleteUser))

	// Clients
	r.GET("/client", ModeWrapper(perqaraHandler.GetClients))
	r.POST("/client", ModeWrapper(perqaraHandler.CreateClient))
	r.PUT("/client", ModeWrapper(perqaraHandler.UpdateClient))
	r.DELETE("/client/:id", ModeWrapper(perqaraHandler.DeleteClient))
}

func AuthenticationSkipper(env string, skipAuthEnv []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var isSkipAuth bool
		for _, skipEnv := range skipAuthEnv {
			if env == skipEnv {
				isSkipAuth = true
				break
			}
		}

		if !isSkipAuth {
			c.Next()
			return
		}

		skipAuthHeader := c.GetHeader("X-SKIP-AUTH")
		if skipAuthHeader == "" {
			c.Next()
			return
		}

		skipAuthTokenHash := os.Getenv("SKIP_AUTH_TOKEN_HASH")
		if err := bcrypt.CompareHashAndPassword([]byte(skipAuthTokenHash), []byte(skipAuthHeader)); err != nil {
			c.Next()
			return
		}

		skipAuthSsoId := os.Getenv("SKIP_AUTH_SSO_ID")
		skipAuthUserRoleId := os.Getenv("SKIP_AUTH_USER_ROLE_ID")
		skipAuthRoleId := os.Getenv("SKIP_AUTH_ROLE_ID")
		skipAuthRoleName := os.Getenv("SKIP_AUTH_ROLE_NAME")
		skipAuthUserName := os.Getenv("SKIP_AUTH_USER_NAME")

		c.Set("skip-auth", struct{}{})
		c.Set("skip-auth-sso-id", skipAuthSsoId)
		c.Set("skip-auth-user-name", skipAuthUserName)
		c.Set("skip-auth-user-role-id", skipAuthUserRoleId)
		c.Set("skip-auth-role-id", skipAuthRoleId)
		c.Set("skip-auth-role-name", skipAuthRoleName)
		c.Next()
	}
}

func ModeWrapper(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if val, ok := c.Get("run-mode"); ok {
			mode, ok := val.(string)
			if ok {
				_, exists := c.Get("public-mode")
				if mode == "public" {
					if exists {
						h(c)
					} else {
						http_response.AbortError(status.Error(codes.PermissionDenied, "api access forbidden"), c)
					}
				} else {
					h(c)
				}
			}
		} else {
			h(c)
		}
	}
}
