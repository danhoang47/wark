package routes

import (
	appcontext "wark/components/app_context"
	"wark/middlewares"
	"wark/modules/priorities/prioritytrans"

	"github.com/gin-gonic/gin"
)

func ConfigPriorityRoutes(rg *gin.RouterGroup, appCtx appcontext.AppContext) {
	priorityRoutes := rg.Group("/priorities")

	priorityRoutes.Use(middlewares.Auth(appCtx))

	priorityRoutes.POST("", prioritytrans.CreatePriority(appCtx))
}
