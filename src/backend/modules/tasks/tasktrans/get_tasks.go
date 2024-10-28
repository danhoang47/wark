package tasktrans

import (
	"net/http"
	"wark/common"
	appcontext "wark/components/app_context"
	"wark/modules/tasks/taskbiz"
	"wark/modules/tasks/taskmodels"
	"wark/modules/tasks/taskrepos"
	"wark/modules/users/usermodels"

	"github.com/gin-gonic/gin"
)

func GetTasks(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		maybeUser, ok := c.Get("user")

		if !ok {
			c.JSON(http.StatusUnauthorized, common.Response{})
			return
		}

		user, ok := maybeUser.(usermodels.User)

		if !ok {
			c.JSON(http.StatusBadRequest, common.Response{})
			return
		}

		getTasksRepo := taskrepos.NewGetTasksRepo(appCtx.GetDB())
		getTasksBiz := taskbiz.NewGetTasksBiz(getTasksRepo)

		tasks := getTasksBiz.GetTasks(user.Id.String(), &taskmodels.GetTaskConds{})

		c.JSON(http.StatusOK, common.Response{
			ErrorCode: 0,
			Data:      tasks,
			Status:    http.StatusOK,
		})
	}
}
