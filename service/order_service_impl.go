package service

import (
	"luwjistik/entity"
	"luwjistik/exception"
	"luwjistik/model"
	"luwjistik/repository"
	"time"

	"github.com/google/uuid"
)

func NewOrderService(orderRepository *repository.OrderRepository) OrderService {
	return &orderServiceImpl{OrderRepository: *orderRepository}
}

type orderServiceImpl struct {
	OrderRepository repository.OrderRepository
}

func (service *orderServiceImpl) Create(order model.CreateOrderRequest) error {
	orderId := uuid.New().String()
	trackingOrder := make([]entity.TrackingOrder, 1)
	trackingOrder[0].Id = uuid.New().String()
	trackingOrder[0].CheckPoints = "Order Created"
	trackingOrder[0].TimeStamp = time.Now()
	trackingOrder[0].OrderId = orderId

	_order := entity.Order{
		Id:               orderId,
		Sender:           order.Sender,
		SenderMobileNo:   order.SenderMobileNo,
		ReceiverAddress:  order.ReceiverAddress,
		ReceiverName:     order.ReceiverName,
		ReceiverMobileNo: order.ReceiverMobileNo,
		Weight:           order.Weight,
		Status:           "1",
	}

	data := entity.OrderDetail{
		Order:          _order,
		TrackingOrders: trackingOrder,
	}

	err := service.OrderRepository.Create(data)
	if err != nil {
		return err
	}
	return nil
}

func (service *orderServiceImpl) GetDetailById(id string) (order model.GetOrderDetailRequest) {
	if id == "" {
		panic(exception.ValidationError{
			Message: "Id cannot be blank",
		})
	}
	_order := service.OrderRepository.GetDetailById(id)

	trackingOrders := []model.GetTrackingOrderRequest{}
	for _, o := range _order.TrackingOrders {
		trackingOrders = append(trackingOrders, model.GetTrackingOrderRequest{
			Id:          o.Id,
			OrderId:     o.OrderId,
			TimeStamp:   o.TimeStamp,
			CheckPoints: o.CheckPoints,
		})
	}
	order.Id = _order.Id
	order.Sender = _order.Sender
	order.SenderMobileNo = _order.Sender
	order.Weight = _order.Weight
	order.Status = _order.Status
	order.ReceiverAddress = _order.ReceiverAddress
	order.ReceiverName = _order.ReceiverName
	order.ReceiverMobileNo = _order.ReceiverMobileNo
	order.TrackingOrders = trackingOrders

	return order
}
