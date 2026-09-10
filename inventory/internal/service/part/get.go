package part

import (
	"boilerplates/inventory/internal/model"
	"context"
)

func (s service) Get(ctx context.Context, uuid string) (model.Part, error) {
	return s.partRepository.Get(ctx, uuid)
}
