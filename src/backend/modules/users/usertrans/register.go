package usertrans

import (
	"log"
	"wark/common"
	appcontext "wark/components/app_context"
	"wark/modules/users/usermodels"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createUser *usermodels.CreateUser

		if c.ShouldBind(&createUser) != nil {
			panic(common.ErrBadRequest)
		} else {
			log.Println(createUser)
		}
	}
}
