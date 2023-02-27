package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/common"
	"simple-rest-api/components"
	"simple-rest-api/modules/restaurant/business"
	"simple-rest-api/modules/restaurant/model"
	"simple-rest-api/modules/restaurant/storage"
)

func HandleCreateOneRestaurant(appContext components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.RestaurantToCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrBadRequest(err, ""))
		}

		store := storage.NewMySQLStore(appContext.GetMainDBConnection())
		biz := business.NewCreateOneRestaurantBiz(store)

		if err := biz.CreateOneRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, common.NewSimpleSuccessResponse(data))
	}
}
