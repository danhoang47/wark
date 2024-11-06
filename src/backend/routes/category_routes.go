package routes

import (
	appcontext "wark/components/app_context"
	"wark/middlewares"
	"wark/modules/categories/categorytrans"

	"github.com/gin-gonic/gin"
)

func ConfigCategoryRoutes(rg *gin.RouterGroup, appCtx appcontext.AppContext) {
	categoryRoutes := rg.Group("/categories")

	categoryRoutes.Use(middlewares.Auth(appCtx))

	categoryRoutes.POST("", categorytrans.CreateCategory(appCtx))
}
