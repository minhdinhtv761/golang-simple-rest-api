package storage

import (
	"context"
	"simple-rest-api/modules/restaurant/model"
)

func (s *mySQLStore) SoftDeleteOneRestaurant(ctx context.Context, id *int) error {
	db := s.db

	if err := db.
		Table(model.Restaurant{}.TableName()).
		Where("id = ?", &id).
		Update("status", "inactive").Error; err != nil {
		return err
	}

	return nil
}
