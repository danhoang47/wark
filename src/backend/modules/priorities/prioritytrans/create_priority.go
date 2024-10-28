package prioritytrans

import (
	"net/http"
	"wark/common"
	appcontext "wark/components/app_context"
	"wark/modules/priorities/prioritybiz"
	"wark/modules/priorities/prioritymodels"
	"wark/modules/priorities/priorityrepos"
	"wark/modules/users/usermodels"

	"github.com/gin-gonic/gin"
)

func CreatePriority(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		maybeUser, exist := c.Get("user")

		if !exist {
			panic(common.ErrUserNotFound)
		}

		user, ok := maybeUser.(usermodels.User)

		if !ok {
			panic(common.ErrUserNotFound)
		}

		var createPriority prioritymodels.CreatePriority

		if err := c.Bind(&createPriority); err != nil {
			c.JSON(http.StatusBadRequest, common.Response{})
		}

		repo := priorityrepos.NewAddPriorityRepo(appCtx.GetDB())
		biz := prioritybiz.NewCreatePriorityBiz(repo)

		rowsAffected, err := biz.CreatePriority(user.Id.String(), &createPriority)

		c.JSON(http.StatusCreated, common.Response{
			ErrorCode: 0,
			Data:      rowsAffected == 1 && err != nil,
		})
	}
}
