package business

import (
	"context"
	"errors"
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
	if data.Name == "" {
		return errors.New("restaurant's name can't be blank")
	}

	err := biz.store.InsertOneRestaurant(ctx, data)

	return err
}
