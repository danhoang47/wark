package middlewares

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"time"

	"wark/common"
	appcontext "wark/components/app_context"
	"wark/components/tokenprovider/jwt"
	"wark/modules/users/usermodels"

	"github.com/gin-gonic/gin"
)

const Expiry = time.Hour * 24 * 7
const JSONRootPath = "$"

func Auth(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")

		if err != nil {
			panic(err)
		}

		jwtProvider := jwt.New(appCtx.GetSecret())
		id, err := jwtProvider.Verify(accessToken)

		if err != nil {
			log.Fatalln(err)
			panic(err)
		}

		db := appCtx.GetDB()
		memCached := appCtx.GetMemCached()

		ctx := context.Background()
		timeOutCtx, cancel := context.WithTimeout(ctx, time.Second*3)
		defer cancel()

		user := usermodels.User{}
		userMemCachedKey := common.GetUserMemCachedKey(id)
		var userJson string

		ch := make(chan string)

		go func() {
			userJson, err := memCached.JSONGet(timeOutCtx, userMemCachedKey, JSONRootPath).Result()

			if err != nil {
				log.Fatalln(err)
			}

			if len(userJson) != 0 {
				ch <- userJson[1 : len(userJson)-1]
			} else {
				ch <- userJson
			}

		}()

		select {
		case <-timeOutCtx.Done():
			panic(common.ErrUserNotFound)
		case userJson = <-ch:
		}

		switch userJson {
		case common.EmptyCachedValue:
			panic(common.ErrUserNotFound)
		case "":
			setCtx := context.Background()
			err = db.Get(&user, `SELECT * FROM users WHERE id = $1`, id)

			if errors.Is(err, sql.ErrNoRows) {
				memCached.JSONSet(setCtx, userMemCachedKey, JSONRootPath, common.EmptyCachedValue)
				panic(common.ErrUserNotFound)
			} else {
				memCached.JSONSet(setCtx, userMemCachedKey, JSONRootPath, &user)
			}

			if ok := memCached.Expire(setCtx, userMemCachedKey, Expiry).Val(); !ok {
				log.Println("cannot set expire for key ", userMemCachedKey)
			}
		default:
			if err := json.Unmarshal([]byte(userJson[1:len(userJson)-1]), &user); err != nil {
				panic(err)
			}
		}

		c.Set("user", user)
		c.Next()
	}
}
