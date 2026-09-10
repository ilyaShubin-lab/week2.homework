package order

import (
	"context"

	"boilerplates/order/internal/model"
	"boilerplates/order/internal/repository/converter"
)

func (r *repository) Create(ctx context.Context, order model.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[order.OrderUUID] = converter.OrderToRepoModel(order)

	return nil
}
