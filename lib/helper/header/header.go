package header

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	customerTransportAppName = "customer-transport"
)

func IsReqFromCustomerTransportApp(c *gin.Context) bool {
	appName := strings.ToLower(strings.TrimSpace(c.Request.Header.Get("X-App-Name")))

	return appName == customerTransportAppName
}
