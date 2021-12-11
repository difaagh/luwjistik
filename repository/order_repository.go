package repository

import "luwjistik/entity"

type OrderRepository interface {
	DeleteAll()
	Create(order entity.OrderDetail) error
	GetById(orderId string) entity.Order
	GetByEmail(email string) entity.Order
	GetDetailById(orderId string) entity.OrderDetail
}
