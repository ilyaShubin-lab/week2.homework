package converter

import (
	"boilerplates/order/internal/model"
	repoModel "boilerplates/order/internal/repository/model"
)

func OrderToRepoModel(order model.Order) repoModel.Order {
	res := repoModel.Order{
		OrderUUID:  order.OrderUUID,
		UserUUID:   order.UserUUID,
		PartUUIDs:  copyStrings(order.PartUUIDs),
		TotalPrice: order.TotalPrice,
		Status:     string(order.Status),
	}

	if order.TransactionUUID != nil {
		txUUID := *order.TransactionUUID
		res.TransactionUUID = &txUUID
	}

	if order.PaymentMethod != nil {
		method := string(*order.PaymentMethod)
		res.PaymentMethod = &method
	}

	return res
}

func OrderToModel(order repoModel.Order) model.Order {
	res := model.Order{
		OrderUUID:  order.OrderUUID,
		UserUUID:   order.UserUUID,
		PartUUIDs:  copyStrings(order.PartUUIDs),
		TotalPrice: order.TotalPrice,
		Status:     model.OrderStatus(order.Status),
	}

	if order.TransactionUUID != nil {
		txUUID := *order.TransactionUUID
		res.TransactionUUID = &txUUID
	}

	if order.PaymentMethod != nil {
		method := model.PaymentMethod(*order.PaymentMethod)
		res.PaymentMethod = &method
	}

	return res
}

func copyStrings(src []string) []string {
	if src == nil {
		return nil
	}

	dst := make([]string, len(src))
	copy(dst, src)

	return dst
}
