package tasktrans

import (
	"errors"
	"io"
	"log"
	"net/http"
	"wark/common"
	appcontext "wark/components/app_context"
	"wark/modules/tasks/taskbiz"
	"wark/modules/tasks/taskmodels"
	"wark/modules/tasks/taskrepos"
	"wark/modules/users/usermodels"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

		var conds taskmodels.GetTaskConds

		if err := c.ShouldBindBodyWithJSON(&conds); err != nil && !errors.Is(err, io.EOF) {
			log.Println("bind GetTaskConds:", err)
			c.JSON(http.StatusBadRequest, common.Response{})
			return
		}

		log.Println(conds.NextCursor)

		getTasksRepo := taskrepos.NewGetTasksRepo(appCtx.GetDB())
		getTasksBiz := taskbiz.NewGetTasksBiz(getTasksRepo)

		tasks := getTasksBiz.GetTasks(user.Id.String(), &conds)

		var nextCursor uuid.NullUUID

		if len(tasks) == 20 {
			nextCursor = uuid.NullUUID{UUID: tasks[19].Id, Valid: true}
		}

		c.JSON(http.StatusOK, common.Response{
			ErrorCode:  0,
			Data:       tasks,
			Status:     http.StatusOK,
			NextCursor: nextCursor,
		})
	}
}
