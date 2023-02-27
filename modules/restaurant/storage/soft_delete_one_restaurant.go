package storage

import (
	"context"
	"simple-rest-api/common"
	"simple-rest-api/modules/restaurant/model"
)

func (s *mySQLStore) SoftDeleteOneRestaurant(ctx context.Context, id *int) error {
	db := s.db

	if err := db.
		Table(model.Restaurant{}.TableName()).
		Where("id = ?", &id).
		Update("status", "inactive").Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
