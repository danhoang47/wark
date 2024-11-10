package routes

import (
	appcontext "wark/components/app_context"
	"wark/middlewares"
	"wark/modules/jobs/jobtrans"

	"github.com/gin-gonic/gin"
)

func ConfigJobRoutes(group *gin.RouterGroup, appCtx appcontext.AppContext) {
	jobRoutes := group.Group("/jobs")

	jobRoutes.Use(middlewares.Auth(appCtx))

	jobRoutes.POST("", jobtrans.CreateJob(appCtx))
}
