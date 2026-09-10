package part

import (
	"boilerplates/inventory/internal/model"
)

const engineUUID = "550e8400-e29b-41d4-a716-446655440001"

func (s *RepositorySuite) TestListEmptyFilterReturnAll() {

	parts, err := s.repo.List(s.ctx, model.PartsFilter{})
	s.Require().NoError(err)
	s.Require().Len(parts, 3)
	//s.Require().NotEmpty(parts)
}

func (s *RepositorySuite) TestListByUUID() {
	allParts, err := s.repo.List(s.ctx, model.PartsFilter{})
	s.Require().NoError(err)
	s.Require().NotEmpty(allParts)

	target := allParts[0]

	parts, err := s.repo.List(s.ctx, model.PartsFilter{UUIDs: []string{target.UUID}})
	s.Require().NoError(err)
	s.Require().Len(parts, 1)
	s.Require().Equal(target.UUID, parts[0].UUID)
}

func (s *RepositorySuite) TestListByName() {
	allParts, err := s.repo.List(s.ctx, model.PartsFilter{})
	s.Require().NoError(err)
	s.Require().NotEmpty(allParts)

	target := allParts[0]

	parts, err := s.repo.List(s.ctx, model.PartsFilter{Names: []string{target.Name}})
	s.Require().NoError(err)
	s.Require().Len(parts, 1)
	s.Require().Equal(target.Name, parts[0].Name)
}

func (s *RepositorySuite) TestListByCategory() {
	allParts, err := s.repo.List(s.ctx, model.PartsFilter{})
	s.Require().NoError(err)
	s.Require().NotEmpty(allParts)

	target := allParts[0]

	parts, err := s.repo.List(s.ctx, model.PartsFilter{Categories: []model.Category{target.Category}})
	s.Require().NoError(err)
	s.Require().NotEmpty(parts)
	for _, p := range parts {
		s.Require().Equal(target.Category, p.Category)
	}
}

func (s *RepositorySuite) TestListByTag() {
	allParts, err := s.repo.List(s.ctx, model.PartsFilter{})
	s.Require().NoError(err)
	s.Require().NotEmpty(allParts)
	var tag string

	for _, p := range allParts {
		if len(p.Tags) > 0 {
			tag = p.Tags[0]
			break
		}
	}
	s.Require().NotEmpty(tag, "в initParts нет ни одной детали с тегами")

	parts, err := s.repo.List(s.ctx, model.PartsFilter{Tags: []string{tag}})
	s.Require().NoError(err)
	s.Require().NotEmpty(parts)
}

func (s *RepositorySuite) TestListByManufacturerCountry() {
	parts, err := s.repo.List(s.ctx, model.PartsFilter{
		ManufacturerCountries: []string{"Россия"},
	})

	s.Require().NoError(err)
	s.Require().Len(parts, 1) // две детали без Manufacturer отсеяны, а не уронили List
	s.Require().Equal(engineUUID, parts[0].UUID)
}
func (s *RepositorySuite) TestListNoMatch() {
	allParts, err := s.repo.List(s.ctx, model.PartsFilter{UUIDs: []string{"nope"}})
	s.Require().NoError(err)
	s.Require().Empty(allParts)
}
