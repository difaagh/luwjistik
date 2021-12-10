package service

import "luwjistik/model"

type OrderService interface {
	GetDetailById(id string) (order model.GetOrderDetailRequest)
	Create(order model.CreateOrderRequest) error
}
