package storage

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/model"
)

func (s *mySQLStore) SelectManyRestaurantsByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Restaurant, error) {
	var result []model.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(model.Restaurant{}.TableName()).Where(conditions)

	if filter != nil {
		if filter.CityId > 0 {
			db = db.Where("city_id = ?", filter.CityId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
