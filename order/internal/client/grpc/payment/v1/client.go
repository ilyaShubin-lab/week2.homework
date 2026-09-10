package v1

import paymentV1 "boilerplates/shared/pkg/proto/payment/v1"

type client struct {
	generatedClient paymentV1.PaymentServiceClient
}

func NewClient(generatedClient paymentV1.PaymentServiceClient) *client {
	return &client{generatedClient: generatedClient}
}
