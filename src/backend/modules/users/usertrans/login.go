package usertrans

import (
	"net/http"
	"wark/common"
	appcontext "wark/components/app_context"
	"wark/components/hasher"
	"wark/modules/users/userbiz"
	"wark/modules/users/usermodels"
	"wark/modules/users/userrepos"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createUser *usermodels.CreateUser

		if c.ShouldBind(&createUser) != nil {
			panic(common.ErrBadRequest)
		}

		hasher := hasher.New()
		addUserRepo := userrepos.NewAddUserRepo(appCtx.GetDB())
		registerUserBiz := userbiz.New(addUserRepo)

		if r, err := registerUserBiz.Register(hasher, createUser); err == nil && r == 1 {
			c.JSON(http.StatusCreated, common.Response{
				ErrorCode: 0,
			})
		} else {
			c.JSON(http.StatusCreated, common.Response{
				ErrorCode: 1,
			})
		}
	}
}
