package business

import "context"

type MarkDeletedOneRestaurantStore interface {
	SoftDeleteOneRestaurant(ctx context.Context, id *int) error
}

type deleteOneRestaurantBiz struct {
	store MarkDeletedOneRestaurantStore
}

func NewDeleteOneRestaurantBiz(store MarkDeletedOneRestaurantStore) *deleteOneRestaurantBiz {
	return &deleteOneRestaurantBiz{store}

}

func (biz *deleteOneRestaurantBiz) DeleteOneRestaurant(ctx context.Context, id *int) error {
	err := biz.store.SoftDeleteOneRestaurant(ctx, id)

	return err
}
