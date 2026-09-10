package v1

import (
	"boilerplates/inventory/internal/service"
	inventoryv1 "boilerplates/shared/pkg/proto/inventory/v1"
)

type api struct {
	inventoryv1.UnimplementedInventoryServiceServer
	partService service.PartService
}

func NewAPI(partService service.PartService) *api {
	return &api{partService: partService}
}
