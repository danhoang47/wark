package middlewares

import (
	"errors"
	"log"
	"net/http"
	"wark/common"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				err, ok := err.(error)

				if !ok {
					log.Fatal(err)
				}

				switch {
				case errors.Is(err, common.ErrBadRequest):
					c.JSON(
						http.StatusBadRequest,
						common.Response{},
					)
				case errors.Is(err, common.ErrUserNotFound):
					c.JSON(
						http.StatusOK,
						common.Response{
							Message: "user not found",
						},
					)
				}
			}
		}()

		c.Next()
	}
}
