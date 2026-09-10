package order

import (
	"boilerplates/order/internal/model"
	"context"
)

func (s *service) Get(ctx context.Context, orderUUID string) (model.Order, error) {
	order, err := s.orderRepository.Get(ctx, orderUUID)
	if err != nil {
		return model.Order{}, err
	}
	return order, nil
}
