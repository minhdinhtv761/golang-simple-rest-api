package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"simple-rest-api/components"
	"simple-rest-api/modules/restaurant/transport"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err := runService(db); err != nil {
		log.Fatal(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD
	appContext := components.NewAppContext(db)

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", transport.HandleCreateOneRestaurant(appContext))
		restaurants.GET("/:id", transport.HandleFindOneRestaurant(appContext))
		restaurants.GET("", transport.HandleFindManyRestaurantsByConditions(appContext))
		restaurants.PATCH("/:id", transport.HandleEditOneRestaurant(appContext))
		restaurants.DELETE("/:id", transport.HandleDeleteOneRestaurant(appContext))
	}

	return r.Run()
}
