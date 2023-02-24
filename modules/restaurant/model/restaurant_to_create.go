package model

type RestaurantToCreate struct {
	Name   string           `json:"name" gorm:"column:name;"`
	Addr   string           `json:"addr" gorm:"column:addr;"`
	Status RestaurantStatus `json:"status,omitempty" gorm:"column:status;type:RestaurantStatus;default:'active'"`
}

func (RestaurantToCreate) TableName() string {
	return Restaurant{}.TableName()
}
