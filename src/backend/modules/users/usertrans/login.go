package usertrans

import (
	"log"
	"net/http"
	"time"
	"wark/common"
	appcontext "wark/components/app_context"
	"wark/components/hasher"
	"wark/components/tokenprovider/jwt"
	"wark/modules/users/userbiz"
	"wark/modules/users/usermodels"
	"wark/modules/users/userrepos"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUser *usermodels.CreateUser

		if c.ShouldBind(&loginUser) != nil {
			panic(common.ErrBadRequest)
		}

		hasher := hasher.New()
		addUserRepo := userrepos.NewGetUserRepo(appCtx.GetDB())
		loginUserBiz := userbiz.NewLoginUserBiz(addUserRepo)

		if user, err := loginUserBiz.Login(hasher, loginUser); err == nil {
			jwtProvider := jwt.New(appCtx.GetSecret())
			token, err := jwtProvider.Generate(user.Id.String(), 15*time.Minute)

			if err != nil {
				log.Fatalln(err)
				c.JSON(http.StatusOK, common.Response{
					ErrorCode: 1,
				})
			}

			c.SetCookie("access_token", token, 15*60, "/", "*", false, false)
			c.JSON(http.StatusOK, common.Response{
				ErrorCode: 0,
				Data:      user,
			})
		} else {
			c.JSON(http.StatusOK, common.Response{
				ErrorCode: 1,
			})
		}
	}
}
