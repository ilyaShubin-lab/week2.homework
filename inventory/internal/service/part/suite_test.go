package part

import (
	"boilerplates/inventory/internal/repository/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite

	ctx context.Context

	partRepository *mocks.MockPartRepository

	service *service
}

func (s *ServiceSuite) SetupTest() {
	s.ctx = context.Background()

	s.partRepository = mocks.NewMockPartRepository(s.T())

	s.service = NewService(
		s.partRepository,
	)
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}
