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
	"strconv"
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

		restaurants.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": err.Error(),
				})

				return
			}

			if err := db.Table(Restaurant{}.TableName()).
				Where("id = ?", id).
				Delete(nil).Error; err != nil {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": err.Error(),
				})

				return
			}

			c.JSON(http.StatusOK, gin.H{"ok": 1})
		})
	}

	return r.Run()
}

//CREATE TABLE `restaurants` (
//`id` int(11) NOT NULL AUTO_INCREMENT,
//`owner_id` int(11) DEFAULT NULL,
//`name` varchar(50) NOT NULL,
//`addr` varchar(255) NOT NULL,
//`city_id` int(11) DEFAULT NULL,
//`lat` double DEFAULT NULL,
//`lng` double DEFAULT NULL,
//`cover` json DEFAULT NULL,
//`logo` json DEFAULT NULL,
//`shipping_fee_per_km` double DEFAULT '0',
//`status` int(11) NOT NULL DEFAULT '1',
//`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//PRIMARY KEY (`id`),
//KEY `owner_id` (`owner_id`) USING BTREE,
//KEY `city_id` (`city_id`) USING BTREE,
//KEY `status` (`status`) USING BTREE
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}
