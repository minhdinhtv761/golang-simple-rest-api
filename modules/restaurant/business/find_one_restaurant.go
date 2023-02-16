package business

import (
	"context"
	"simple-rest-api/modules/restaurant/model"
)

type SelectOneRestaurantStore interface {
	SelectOneRestaurant(ctx context.Context, id *int, data *model.Restaurant) error
}

type findOneRestaurantBiz struct {
	store SelectOneRestaurantStore
}

func NewFindOneRestaurantBiz(store SelectOneRestaurantStore) *findOneRestaurantBiz {
	return &findOneRestaurantBiz{store}
}

func (biz *findOneRestaurantBiz) FindOneRestaurant(ctx context.Context, id *int, data *model.Restaurant) error {
	err := biz.store.SelectOneRestaurant(ctx, id, data)

	return err
}
