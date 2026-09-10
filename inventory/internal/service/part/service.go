package part

import "boilerplates/inventory/internal/repository"

type service struct {
	partRepository repository.PartRepository
}

func NewService(partRepository repository.PartRepository) *service {
	return &service{partRepository: partRepository}
}
