package business

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/model"
)

type SelectManyRestaurantsByConditionsStore interface {
	SelectManyRestaurantsByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		paging *common.Paging,
		filter *model.RestaurantFilter,
		sort string,
		moreKeys ...string,
	) ([]model.Restaurant, error)
}

type findManyRestaurantsByConditionsBiz struct {
	store SelectManyRestaurantsByConditionsStore
}

func NewFindManyRestaurantsByConditionsBiz(
	store SelectManyRestaurantsByConditionsStore,
) *findManyRestaurantsByConditionsBiz {
	return &findManyRestaurantsByConditionsBiz{store}
}

func (biz *findManyRestaurantsByConditionsBiz) FindManyRestaurantsByConditions(
	ctx context.Context,
	paging *common.Paging,
	filter *model.RestaurantFilter,
	sort string,
) ([]model.Restaurant, error) {
	result, err := biz.store.SelectManyRestaurantsByConditions(ctx, nil, paging, filter, sort)

	return result, err
}
