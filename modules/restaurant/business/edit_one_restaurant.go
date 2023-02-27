package business

import (
	"context"
	"simple-rest-api/common"
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
		if err == common.ErrRecordNotFound {
			return common.ErrCannotUpdateEntityResource(model.EntityName, err)
		}

		return common.ErrCannotUpdateEntityResource(model.EntityName, err)
	}

	if oldData.Status == model.Inactive {
		return common.ErrCannotUpdateEntityResource(model.EntityName, err)
	}

	if err := biz.store.UpdateOneRestaurant(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntityResource(model.EntityName, err)
	}

	return nil
}
