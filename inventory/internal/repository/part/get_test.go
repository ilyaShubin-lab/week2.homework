package part

import "boilerplates/inventory/internal/model"

const (
	existingPartUUID = "550e8400-e29b-41d4-a716-446655440002"
	existingPartName = "Топливный бак ТБ-200"
)

func (s *RepositorySuite) TestGetSuccess() {
	part, err := s.repo.Get(s.ctx, existingPartUUID)
	s.Require().NoError(err)
	s.Require().Equal(existingPartUUID, part.UUID)
	s.Require().Equal(existingPartName, part.Name)
	s.Require().NotZero(part.Price)
	s.Require().NotZero(part.UpdatedAt)
}

func (s *RepositorySuite) TestGetNotFound() {
	part, err := s.repo.Get(s.ctx, "unknown UUID")
	//s.Require().NotEqual("unknown UUID", part.UUID)
	s.Require().ErrorIs(err, model.ErrPartNotFound)
	s.Require().Equal(model.Part{}, part)
}
