package business

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/model"
)

type InsertOneRestaurantStore interface {
	InsertOneRestaurant(ctx context.Context, data *model.RestaurantToCreate) error
}

type createOneRestaurantBiz struct {
	store InsertOneRestaurantStore
}

func NewCreateOneRestaurantBiz(store InsertOneRestaurantStore) *createOneRestaurantBiz {
	return &createOneRestaurantBiz{store}
}

func (biz *createOneRestaurantBiz) CreateOneRestaurant(ctx context.Context, data *model.RestaurantToCreate) error {
	if err := data.Validate(); err != nil {
		return common.ErrBadRequest(err, err.Error())
	}

	err := biz.store.InsertOneRestaurant(ctx, data)

	return common.ErrCannotCreateEntityResource(model.EntityName, err)
}
