package business

import (
	"context"
	"errors"
	"simple-rest-api/modules/restaurant/model"
)

type UpdateOneRestaurantStore interface {
	SelectOneRestaurantByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*model.Restaurant, error)
	UpdateOneRestaurant(ctx context.Context, id *int, data *model.RestaurantToUpdate) error
}

type editOneRestaurantBiz struct {
	store UpdateOneRestaurantStore
}

func NewEditOneRestaurantBiz(store UpdateOneRestaurantStore) *editOneRestaurantBiz {
	return &editOneRestaurantBiz{store}
}

func (biz editOneRestaurantBiz) EditOneRestaurant(ctx context.Context, id *int, data *model.RestaurantToUpdate) error {
	oldData, err := biz.store.SelectOneRestaurantByConditions(
		ctx,
		map[string]interface{}{"id": id},
	)

	if err != nil {
		return err
	}

	if oldData.Status == "inactive" {
		return errors.New("restaurant has already been deleted")
	}

	if err := biz.store.UpdateOneRestaurant(ctx, id, data); err != nil {
		return err
	}

	return nil
}
