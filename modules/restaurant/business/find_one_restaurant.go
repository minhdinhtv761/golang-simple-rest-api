package business

import (
	"context"
	"errors"
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
		return nil, err
	}

	if data.Status == "inactive" {
		return nil, errors.New("restaurant has already been deleted")
	}

	return data, err
}
