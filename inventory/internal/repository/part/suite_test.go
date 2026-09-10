package part

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RepositorySuite struct {
	suite.Suite
	ctx  context.Context
	repo *repository
}

func (s *RepositorySuite) SetupTest() {
	s.ctx = context.Background()
	s.repo = NewRepository() // настоящее хранилище, никаких моков
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}
