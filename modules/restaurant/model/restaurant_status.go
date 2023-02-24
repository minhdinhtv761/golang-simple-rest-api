package model

type RestaurantStatus string

const (
	Active   RestaurantStatus = "active"
	Inactive RestaurantStatus = "inactive"
)

func (r RestaurantStatus) GormDataType() string {
	return "ENUM('active','inactive')"
}
