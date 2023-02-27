package model

import (
	"fmt"
	"simple-rest-api/common"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string           `json:"name" gorm:"column:name;"`
	Addr            string           `json:"addr" gorm:"column:addr;"`
	Status          RestaurantStatus `json:"status" gorm:"column:status;type:RestaurantStatus;'"`
}

func (Restaurant) TableName() string {
	return fmt.Sprintf("%ss", EntityName)
}
