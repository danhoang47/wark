package routes

import (
	appcontext "wark/components/app_context"
	"wark/middlewares"
	"wark/modules/tasks/tasktrans"

	"github.com/gin-gonic/gin"
)

func ConfigTasksRoutes(c *gin.RouterGroup, appCtx appcontext.AppContext) {
	tasksRoutes := c.Group("/tasks")

	tasksRoutes.Use(middlewares.Auth(appCtx))

	tasksRoutes.GET("", tasktrans.GetTasks(appCtx))
}
