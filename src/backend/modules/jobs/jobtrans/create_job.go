package jobtrans

import (
	"log"
	"net/http"
	"wark/common"
	appcontext "wark/components/app_context"
	"wark/modules/jobs/jobmodels"

	"github.com/gin-gonic/gin"
)

func CreateJob(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var job jobmodels.CreateJob

		if err := c.ShouldBindBodyWithJSON(&job); err != nil {
			panic(common.ErrBadRequest)
		}

		log.Println(job)

		c.JSON(http.StatusCreated, common.Response{})
	}
}
