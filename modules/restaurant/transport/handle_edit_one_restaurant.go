package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/common"
	"simple-rest-api/components"
	"simple-rest-api/modules/restaurant/business"
	"simple-rest-api/modules/restaurant/model"
	"simple-rest-api/modules/restaurant/storage"
	"strconv"
)

func HandleEditOneRestaurant(appContext components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err, ""))
		}

		var data model.RestaurantToUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrBadRequest(err, ""))
		}

		store := storage.NewMySQLStore(appContext.GetMainDBConnection())
		biz := business.NewEditOneRestaurantBiz(store)

		if err := biz.EditOneRestaurant(c.Request.Context(), &id, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewDoneSuccessResponse())
	}
}
