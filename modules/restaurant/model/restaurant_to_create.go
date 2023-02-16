package model

type RestaurantToCreate struct {
	Id   int    `json:"id,omitempty" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantToCreate) TableName() string {
	return Restaurant{}.TableName()
}
