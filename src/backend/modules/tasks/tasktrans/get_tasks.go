package tasktrans

import (
	"log"
	"net/http"
	appcontext "wark/components/app_context"

	"github.com/gin-gonic/gin"
)

func GetTasks(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println(c.Get("user"))

		c.JSON(http.StatusOK, map[string]interface{}{})
	}
}
