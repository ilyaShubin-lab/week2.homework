package order

import (
	grpcClient "boilerplates/order/internal/client/grpc"
	"boilerplates/order/internal/repository"
)

type service struct {
	orderRepository repository.OrderRepository
	inventoryClient grpcClient.InventoryClient
	paymentClient   grpcClient.PaymentClient
}

func NewService(
	orderRepository repository.OrderRepository,
	inventoryClient grpcClient.InventoryClient,
	paymentClient grpcClient.PaymentClient,
) *service {
	return &service{orderRepository, inventoryClient, paymentClient}
}
