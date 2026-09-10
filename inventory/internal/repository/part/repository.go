package part

// имя папки

import (
	def "boilerplates/inventory/internal/repository"
	repoModel "boilerplates/inventory/internal/repository/model"
	"sync"
)

// Компилятор проверит, что все методы интерфейса на месте
var _ def.PartRepository = (*repository)(nil)

type repository struct {
	mu   sync.RWMutex
	data map[string]repoModel.Part
}

func NewRepository() *repository {

	repo := &repository{data: make(map[string]repoModel.Part)}
	repo.initParts() // тестовые данные
	return repo
}
