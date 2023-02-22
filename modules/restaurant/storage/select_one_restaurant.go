package storage

import (
	"context"
	"simple-rest-api/modules/restaurant/model"
)

func (s *mySQLStore) SelectOneRestaurant(ctx context.Context, id *int, data *model.Restaurant) error {
	db := s.db

	if err := db.Where("id = ?", &id).First(data).Error; err != nil {
		return err
	}

	return nil
}
