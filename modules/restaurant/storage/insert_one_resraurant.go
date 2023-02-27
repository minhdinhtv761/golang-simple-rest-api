package storage

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/model"
)

func (s *mySQLStore) InsertOneRestaurant(ctx context.Context, data *model.RestaurantToCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
