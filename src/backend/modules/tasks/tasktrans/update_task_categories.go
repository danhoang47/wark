package tasktrans

import (
	"log"
	"net/http"
	"wark/common"
	appcontext "wark/components/app_context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UpdateTaskCategories(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskId := c.Param("id")
		var updatedCategories []uuid.UUID

		if err := c.Bind(&updatedCategories); err != nil {
			log.Println(err)
			panic(common.ErrBadRequest)
		}

		log.Println(taskId, updatedCategories)

		c.JSON(http.StatusOK, common.Response{})
	}
}
