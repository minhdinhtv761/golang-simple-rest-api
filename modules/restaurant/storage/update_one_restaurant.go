package storage

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/model"
)

func (s *mySQLStore) UpdateOneRestaurant(ctx context.Context, id *int, data *model.RestaurantToUpdate) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
