package service

import "luwjistik/model"

type OrderService interface {
	GetDetailById(id, email string) (order model.GetOrderDetailRequest)
	Create(order model.CreateOrderRequest) error
}
