package model

type RestaurantToUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantToUpdate) TableName() string {
	return Restaurant{}.TableName()
}
