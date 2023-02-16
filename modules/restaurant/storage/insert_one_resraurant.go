package storage

import (
	"context"
	"simple-rest-api/modules/restaurant/model"
)

func (s *mySQLStore) InsertOneRestaurant(ctx context.Context, data *model.RestaurantToCreate) error {
	db := s.db

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
