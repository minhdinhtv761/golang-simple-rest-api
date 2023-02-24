package common

import "reflect"

type Sort struct {
	SortBy    string `json:"sort_by,omitempty" form:"sort_by"`
	SortOrder string `json:"sort_order,omitempty" form:"sort_order"`
}

func (sort *Sort) IsEmpty() bool {
	return reflect.ValueOf(*sort).IsZero()
}
