package order

import (
	repoModel "boilerplates/order/internal/repository/model"
	"sync"
)

type repository struct {
	mu   sync.RWMutex
	data map[string]repoModel.Order
}

func NewRepository() *repository {
	return &repository{data: make(map[string]repoModel.Order)}
}
