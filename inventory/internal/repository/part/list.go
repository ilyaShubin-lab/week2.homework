package part

import (
	"boilerplates/inventory/internal/model"
	"boilerplates/inventory/internal/repository/converter"
	repoModel "boilerplates/inventory/internal/repository/model"
	"context"
	"slices"
)

func (r *repository) List(ctx context.Context, filter model.PartsFilter) ([]model.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	res := make([]model.Part, 0, len(r.data))
	for _, part := range r.data {
		if !matches(part, filter) {
			continue
		}
		res = append(res, converter.PartToModel(part))
	}
	return res, nil
}

func matches(part repoModel.Part, filter model.PartsFilter) bool {

	if len(filter.UUIDs) > 0 && !slices.Contains(filter.UUIDs, part.UUID) {
		return false
	}

	if len(filter.Names) > 0 && !slices.Contains(filter.Names, part.Name) {
		return false
	}

	if len(filter.Categories) > 0 && !slices.Contains(filter.Categories, model.Category(part.Category)) {
		return false
	}

	if len(filter.ManufacturerCountries) > 0 && (part.Manufacturer == nil || !slices.Contains(filter.ManufacturerCountries, part.Manufacturer.Country)) {
		return false
	}

	if len(filter.Tags) > 0 && !containTags(part.Tags, filter.Tags) {
		return false
	}

	return true
}

func containTags(partTags, filterTags []string) bool {
	for _, v := range partTags {
		if slices.Contains(filterTags, v) {
			return true
		}

	}
	return false

}
