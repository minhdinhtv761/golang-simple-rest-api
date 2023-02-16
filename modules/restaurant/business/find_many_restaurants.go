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
		filter *model.Filter,
		paging *common.Paging,
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
	filter *model.Filter,
	paging *common.Paging,
) ([]model.Restaurant, error) {
	result, err := biz.store.SelectManyRestaurantsByConditions(ctx, nil, filter, paging)

	return result, err
}
