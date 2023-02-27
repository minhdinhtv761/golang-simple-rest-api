package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/common"
	"simple-rest-api/components"
	"simple-rest-api/modules/restaurant/business"
	"simple-rest-api/modules/restaurant/storage"
	"strconv"
)

func HandleFindOneRestaurant(appContext components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrBadRequest(err, ""))
		}

		store := storage.NewMySQLStore(appContext.GetMainDBConnection())
		biz := business.NewFindOneRestaurantBiz(store)

		data, err := biz.FindOneRestaurant(c.Request.Context(), &id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSimpleSuccessResponse(data))
	}
}
