package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	inventoryV1API "boilerplates/inventory/internal/api/inventory/v1"
	partRepository "boilerplates/inventory/internal/repository/part"
	partService "boilerplates/inventory/internal/service/part"
	inventoryV1 "boilerplates/shared/pkg/proto/inventory/v1"
)

const grpcPort = "50051"

func main() {
	// --- хранилище
	repo := partRepository.NewRepository()
	// --- бизнес-логика
	svc := partService.NewService(repo)
	// --- транспорт
	apiV1 := inventoryV1API.NewAPI(svc)
	// --- сеть
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	s := grpc.NewServer()

	// вторая проверка соответствия контракту
	inventoryV1.RegisterInventoryServiceServer(s, apiV1)

	// Чтобы grpcurl / Postman видели список методов без .proto
	reflection.Register(s)

	log.Printf("gRPC server listening on :%s", grpcPort)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("serve: %v", err)
	}
}
