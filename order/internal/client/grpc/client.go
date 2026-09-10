package grpc

import (
	"boilerplates/order/internal/model"
	"context"
)

type InventoryClient interface {
	ListParts(ctx context.Context, filter model.PartsFilter) ([]model.Part, error)
}

type PaymentClient interface {
	PayOrder(ctx context.Context, orderUUID, userUUID string, method model.PaymentMethod) (string, error)
}
