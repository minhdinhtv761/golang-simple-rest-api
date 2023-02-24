package business

import (
	"context"
	"errors"
	"simple-rest-api/modules/restaurant/model"
)

type SoftDeleteOneRestaurantStore interface {
	SelectOneRestaurantByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*model.Restaurant, error)
	SoftDeleteOneRestaurant(ctx context.Context, id *int) error
}

type deleteOneRestaurantBiz struct {
	store SoftDeleteOneRestaurantStore
}

func NewDeleteOneRestaurantBiz(store SoftDeleteOneRestaurantStore) *deleteOneRestaurantBiz {
	return &deleteOneRestaurantBiz{store}

}

func (biz *deleteOneRestaurantBiz) DeleteOneRestaurant(ctx context.Context, id *int) error {
	oldData, err := biz.store.SelectOneRestaurantByConditions(
		ctx,
		map[string]interface{}{
			"id": id,
		},
	)

	if err != nil {
		return err
	}

	if oldData.Status == "inactive" {
		return errors.New("restaurant has already been deleted")
	}

	if err := biz.store.SoftDeleteOneRestaurant(ctx, id); err != nil {
		return err
	}

	return nil
}
