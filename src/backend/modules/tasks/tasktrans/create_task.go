package tasktrans

import (
	"log"
	"net/http"
	"wark/common"
	appcontext "wark/components/app_context"
	"wark/modules/tasks/taskbiz"
	"wark/modules/tasks/taskmodels"
	"wark/modules/tasks/taskrepos"
	"wark/modules/users/usermodels"

	"github.com/gin-gonic/gin"
)

func CreateTask(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task *taskmodels.CreateTask

		if err := c.Bind(&task); err != nil {
			panic(err)
		}

		user, ok := c.MustGet("user").(*usermodels.User)

		log.Println(user)

		if !ok {
			panic(common.ErrUserNotFound)
		}

		task.CreatorId = user.Id

		addTaskRepo := taskrepos.NewAddTaskRepo(appCtx.GetDB())
		addTaskBiz := taskbiz.NewAddTaskRepo(addTaskRepo)

		ok = addTaskBiz.AddTask(task)

		if ok {
			c.JSON(http.StatusCreated, common.Response{
				Status:  0,
				Message: "add task successfully",
			})
		}
	}
}
