package categorytrans

import (
	"net/http"
	"wark/common"
	appcontext "wark/components/app_context"
	"wark/modules/categories/categorybiz"
	"wark/modules/categories/categorymodels"
	"wark/modules/categories/categoryrepos"
	"wark/modules/users/usermodels"

	"github.com/gin-gonic/gin"
)

func CreateCategory(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		maybeUser, exist := c.Get("user")

		if !exist {
			panic(common.ErrUserNotFound)
		}

		user, ok := maybeUser.(usermodels.User)

		if !ok {
			panic(common.ErrUserNotFound)
		}

		var createCategory categorymodels.CreateCategory

		if err := c.Bind(&createCategory); err != nil {
			c.JSON(http.StatusBadRequest, common.Response{})
		}

		repo := categoryrepos.NewAddCategoryRepo(appCtx.GetDB())
		biz := categorybiz.NewCreateCategoryBiz(repo)

		rowsAffected, err := biz.CreateCategory(user.Id.String(), &createCategory)

		c.JSON(http.StatusCreated, common.Response{
			ErrorCode: 0,
			Data:      rowsAffected == 1 && err != nil,
		})
	}
}
