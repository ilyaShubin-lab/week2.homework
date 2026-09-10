package converter

import (
	"boilerplates/order/internal/model"
	inventoryV1 "boilerplates/shared/pkg/proto/inventory/v1"
	paymentV1 "boilerplates/shared/pkg/proto/payment/v1"
)

func PartsFilterToProto(filter model.PartsFilter) *inventoryV1.PartsFilter {
	return &inventoryV1.PartsFilter{
		Uuids: filter.UUIDs,
	}
}

func PartToModel(part *inventoryV1.Part) model.Part {
	return model.Part{
		UUID:  part.GetUuid(),
		Price: part.GetPrice(),
	}
}

func PartListToModel(parts []*inventoryV1.Part) []model.Part {
	res := make([]model.Part, 0, len(parts))

	for _, part := range parts {
		res = append(res, PartToModel(part))
	}

	return res
}

func PaymentMethodToProto(method model.PaymentMethod) paymentV1.PaymentMethod {
	switch method {
	case model.PaymentMethodCard:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_CARD
	case model.PaymentMethodSBP:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_SBP
	case model.PaymentMethodCreditCard:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD
	case model.PaymentMethodInvestorMoney:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY
	default:
		return paymentV1.PaymentMethod_PAYMENT_METHOD_UNSPECIFIED
	}
}
