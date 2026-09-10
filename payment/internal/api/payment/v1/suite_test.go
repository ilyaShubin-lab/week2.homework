package v1

import (
	"boilerplates/payment/internal/service/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type APISuite struct {
	suite.Suite
	ctx            context.Context
	paymentService *mocks.MockPaymentService
	api            *api
}

func (s *APISuite) SetupTest() {

	s.ctx = context.Background()
	s.paymentService = mocks.NewMockPaymentService(s.T())
	s.api = NewAPI(s.paymentService)
}

func TestServiceSuite(t *testing.T) { suite.Run(t, new(APISuite)) }
