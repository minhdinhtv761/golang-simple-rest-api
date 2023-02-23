package storage

import (
	"context"
	"simple-rest-api/modules/restaurant/model"
)

func (s *mySQLStore) SelectOneRestaurantByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*model.Restaurant, error) {
	var data model.Restaurant

	db := s.db

	if err := db.Where(conditions).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
