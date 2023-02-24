package model

import "reflect"

type RestaurantFilter struct {
	CityId int `json:"city_id,omitempty" form:"city_id"`
}

func (filter RestaurantFilter) IsEmpty() bool {
	return reflect.ValueOf(filter).IsZero()
}
