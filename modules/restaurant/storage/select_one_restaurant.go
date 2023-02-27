package storage

import (
	"context"
	"gorm.io/gorm"
	"simple-rest-api/common"
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
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
