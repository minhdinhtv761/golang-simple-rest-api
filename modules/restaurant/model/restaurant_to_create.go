package model

import (
	"errors"
	"strings"
)

type RestaurantToCreate struct {
	ID     int              `json:"id" gorm:"column:id"`
	Name   string           `json:"name" gorm:"column:name;"`
	Addr   string           `json:"addr" gorm:"column:addr;"`
	Status RestaurantStatus `json:"status,omitempty" gorm:"column:status;type:RestaurantStatus;default:'active'"`
}

func (*RestaurantToCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (r *RestaurantToCreate) Validate() error {
	r.Name = strings.TrimSpace(r.Name)
	r.Addr = strings.TrimSpace(r.Addr)

	if len(r.Addr) == 0 {
		return errors.New("restaurant's address can't be blank")
	}

	if len(r.Addr) == 0 {
		return errors.New("restaurant's address can't be blank")
	}

	return nil
}
