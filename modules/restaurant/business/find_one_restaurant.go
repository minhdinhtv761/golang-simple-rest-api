package business

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/model"
)

type SelectOneRestaurantStore interface {
	SelectOneRestaurantByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*model.Restaurant, error)
}

type findOneRestaurantBiz struct {
	store SelectOneRestaurantStore
}

func NewFindOneRestaurantBiz(store SelectOneRestaurantStore) *findOneRestaurantBiz {
	return &findOneRestaurantBiz{store}
}

func (biz *findOneRestaurantBiz) FindOneRestaurant(
	ctx context.Context,
	id *int,
) (*model.Restaurant, error) {
	data, err := biz.store.SelectOneRestaurantByConditions(ctx, map[string]interface{}{"id": *id})

	if err != nil {
		if err == common.ErrRecordNotFound {
			return nil, common.ErrCannotGetEntityResource(model.EntityName, err)
		}

		return nil, common.ErrCannotGetEntityResource(model.EntityName, err)
	}

	if data.Status == model.Inactive {
		return nil, common.ErrCannotGetEntityResource(model.EntityName, err)
	}

	return data, err
}
