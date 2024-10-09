package middlewares

import (
	"context"
	"database/sql"
	"errors"

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

		if err != nil {
			panic(err)
		}

		var user *usermodels.User
		context := context.Background()
		memCached := appCtx.GetMemCached()
		db := appCtx.GetDB()
		userMemCachedKey := common.GetUserMemCachedKey(id)
		userJson, err := memCached.JSONGet(context, userMemCachedKey, "$").Result()

		if userJson == common.EmptyCachedValue {
			panic(common.ErrUserNotFound)
		}

		if err != nil {
			err := db.Get(user, "SELECT * FROM USER WHERE id = $1", id)

			if errors.Is(err, sql.ErrNoRows) {
				memCached.JSONSet(context, userMemCachedKey, "$", common.EmptyCachedValue)
				panic(common.ErrUserNotFound)
			} else {
				memCached.JSONSet(context, userMemCachedKey, "$", user)
			}
		}

		c.Set("user", user)
		c.Next()
	}
}
