package middlewares

import "github.com/gin-gonic/gin"

type Authenticator interface {
	Authenticate()
}

func Auth(c *gin.Context, provider Authenticator) {

}
