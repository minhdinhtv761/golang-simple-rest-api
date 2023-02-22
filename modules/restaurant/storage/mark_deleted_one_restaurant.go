package storage

import (
	"context"
	"simple-rest-api/modules/restaurant/model"
)

func (s *mySQLStore) MarkDeletedOneRestaurant(ctx context.Context, id *int) error {
	db := s.db

	if err := db.
		Table(model.Restaurant{}.TableName()).
		Where("id = ? AND status = ?", &id, 1).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
