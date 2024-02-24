package http_handler

import (
	"context"

	"perqara/config"
	perqara "perqara/domain/perqara"

	"github.com/gin-gonic/gin"
)

type handler struct {
	ctx     context.Context
	usecase perqara.PerqaraUsecaseInterface
	config  config.Config
}

type PerqaraHandler interface {
	// Login
	Login(c *gin.Context)

	// ContentType
	GetContentTypes(c *gin.Context)
	CreateContentType(c *gin.Context)

	// Permission
	GetPermissionsDropdown(c *gin.Context)

	// Role
	GetRoles(c *gin.Context)
	GetRoleDetail(c *gin.Context)
	GetTotalRoles(c *gin.Context)
	GetRolesDropdown(c *gin.Context)
	CreateRole(c *gin.Context)
	UpdateRole(c *gin.Context)
	DeleteRole(c *gin.Context)

	// User
	GetUsers(c *gin.Context)
	GetUserDetail(c *gin.Context)
	GetTotalUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)

	// Client
	GetClients(c *gin.Context)
	CreateClient(c *gin.Context)
	UpdateClient(c *gin.Context)
	DeleteClient(c *gin.Context)
}

func NewPerqaraHandler(
	ctx context.Context,
	usecase perqara.PerqaraUsecaseInterface,
	cfg config.Config,
) PerqaraHandler {
	return &handler{
		ctx:     ctx,
		usecase: usecase,
		config:  cfg,
	}
}
