package business

import (
	"context"
	"errors"
	"simple-rest-api/modules/restaurant/model"
)

type UpdateOneRestaurantStore interface {
	UpdateOneRestaurant(ctx context.Context, id *int, data *model.RestaurantToUpdate) error
}

type editOneRestaurantBiz struct {
	store UpdateOneRestaurantStore
}

func NewEditOneRestaurantBiz(store UpdateOneRestaurantStore) *editOneRestaurantBiz {
	return &editOneRestaurantBiz{store}
}

func (biz editOneRestaurantBiz) EditOneRestaurant(ctx context.Context, id *int, data *model.RestaurantToUpdate) error {
	if name := data.Name; *name == "" {
		return errors.New("restaurant's name can't be blank")
	}

	err := biz.store.UpdateOneRestaurant(ctx, id, data)

	return err
}
