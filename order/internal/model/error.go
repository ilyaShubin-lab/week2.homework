package model

import "errors"

var (
	ErrOrderNotFound    = errors.New("order not found")
	ErrOrderAlreadyPaid = errors.New("order already paid")
	ErrOrderCancelled   = errors.New("order is cancelled")

	ErrPartsNotFound        = errors.New("some parts not found")
	ErrInventoryUnavailable = errors.New("inventory service unavailable")
	ErrPaymentUnavailable   = errors.New("payment service unavailable")
)
