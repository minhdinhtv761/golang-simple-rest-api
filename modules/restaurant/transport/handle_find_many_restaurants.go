package transport

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/common"
	"simple-rest-api/components"
	"simple-rest-api/modules/restaurant/business"
	"simple-rest-api/modules/restaurant/model"
	"simple-rest-api/modules/restaurant/storage"
)

func mapSortParams(sort common.Sort) string {
	var sortBy, sortOrder string
	switch sort.SortBy {
	default:
		sortBy = "id"
	}

	if sort.SortOrder == "asc" {
		sortOrder = "asc"
	} else {
		sortOrder = "desc"
	}

	return fmt.Sprintf("%s %s", sortBy, sortOrder)
}

func HandleFindManyRestaurantsByConditions(appContext components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter model.RestaurantFilter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrBadRequest(err, ""))
		}

		var sort common.Sort

		if err := c.ShouldBind(&sort); err != nil {
			panic(common.ErrBadRequest(err, ""))
		}

		var paging common.Paging

		paging.Fulfill()

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrBadRequest(err, ""))
		}

		store := storage.NewMySQLStore(appContext.GetMainDBConnection())
		biz := business.NewFindManyRestaurantsByConditionsBiz(store)

		result, err := biz.FindManyRestaurantsByConditions(
			c.Request.Context(),
			&paging,
			&filter,
			mapSortParams(sort),
		)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(
			result,
			&paging,
			filter,
			&sort,
		))
	}
}
