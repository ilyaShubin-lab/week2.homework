package order

import (
	"boilerplates/order/internal/model"
	"boilerplates/order/internal/repository/converter"

	//"boilerplates/order/internal/repository/model"
	"context"
)

func (r *repository) Get(ctx context.Context, orderUUID string) (model.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, ok := r.data[orderUUID]
	if !ok {
		return model.Order{}, model.ErrOrderNotFound
	}

	return converter.OrderToModel(order), nil
}
