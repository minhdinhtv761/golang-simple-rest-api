package transport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/components"
	"simple-rest-api/modules/restaurant/business"
	"simple-rest-api/modules/restaurant/storage"
	"strconv"
)

func HandleDeleteOneRestaurant(appContext components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewMySQLStore(appContext.GetMainDBConnection())
		biz := business.NewDeleteOneRestaurantBiz(store)

		if err := biz.DeleteOneRestaurant(c.Request.Context(), &id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusNoContent, gin.H{})
	}
}