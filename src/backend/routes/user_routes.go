package routes

import (
	appcontext "wark/components/app_context"
	"wark/modules/users/usertrans"

	"github.com/gin-gonic/gin"
)

func ConfigUserRoutes(c *gin.RouterGroup, appCtx appcontext.AppContext) {
	userGroup := c.Group("/users")

	userGroup.POST("/register", usertrans.Register(appCtx))
}
