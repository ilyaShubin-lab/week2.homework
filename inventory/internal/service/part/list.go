package part

import (
	"boilerplates/inventory/internal/model"
	"context"
)

func (s *service) List(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	return s.partRepository.List(ctx, filter)
}
