package business

import (
	"context"
	"simple-rest-api/common"
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
		if err == common.ErrRecordNotFound {
			return common.ErrCannotDeleteEntityResource(model.EntityName, err)
		}

		return common.ErrCannotDeleteEntityResource(model.EntityName, err)
	}

	if oldData.Status == model.Inactive {
		return common.ErrCannotDeleteEntityResource(model.EntityName, err)
	}

	if err := biz.store.SoftDeleteOneRestaurant(ctx, id); err != nil {
		return common.ErrCannotDeleteEntityResource(model.EntityName, err)
	}

	return nil
}
