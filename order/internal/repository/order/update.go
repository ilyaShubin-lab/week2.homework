package order

import (
	"boilerplates/order/internal/model"
	"boilerplates/order/internal/repository/converter"

	//"boilerplates/order/internal/repository/model"
	"context"
)

func (r *repository) Update(ctx context.Context, order model.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.data[order.OrderUUID]; !ok {
		return model.ErrOrderNotFound
	}

	r.data[order.OrderUUID] = converter.OrderToRepoModel(order)

	return nil
}
