package middlewares

import (
	"log"

	"wark/common"
	appcontext "wark/components/app_context"
	"wark/components/tokenprovider/jwt"
	"wark/modules/users/usermodels"

	"github.com/gin-gonic/gin"
)

func Auth(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")

		if err != nil {
			panic(err)
		}

		jwtProvider := jwt.New(appCtx.GetSecret())
		id, err := jwtProvider.Verify(accessToken)

		log.Println(id)

		if err != nil {
			log.Fatalln(err)
			panic(err)
		}

		user := &usermodels.User{}
		db := appCtx.GetDB()
		// userMemCachedKey := common.GetUserMemCachedKey(id)
		// userJson, err := memCached.JSONGet(context, userMemCachedKey, "$").Result()

		// if userJson == common.EmptyCachedValue {
		// 	panic(common.ErrUserNotFound)
		// }

		if err := db.Get(user, `SELECT * FROM users WHERE id = $1`, id); err != nil {
			c.Abort()
			panic(common.ErrUserNotFound)
		}

		log.Println(user)

		c.Set("user", user)
		c.Next()
	}
}
