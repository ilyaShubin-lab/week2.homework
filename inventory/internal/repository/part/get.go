package part

import (
	"boilerplates/inventory/internal/model"
	"boilerplates/inventory/internal/repository/converter"
	"context"
)

func (r *repository) Get(ctx context.Context, uuid string) (model.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	part, ok := r.data[uuid]
	if !ok {
		return model.Part{}, model.ErrPartNotFound
	}

	return converter.PartToModel(part), nil

}
