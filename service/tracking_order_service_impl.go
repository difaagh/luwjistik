package service

import (
	"errors"
	"luwjistik/entity"
	"luwjistik/model"
	"luwjistik/repository"
	"time"

	"github.com/google/uuid"
)

func NewTrackingOrderService(trackingOrderRepository *repository.TrackingOrderRepository, orderRepository *repository.OrderRepository) TrackingOrderService {
	return &trackingOrderServiceImpl{TrackingOrderRepository: *trackingOrderRepository, OrderRepository: *orderRepository}
}

type trackingOrderServiceImpl struct {
	TrackingOrderRepository repository.TrackingOrderRepository
	OrderRepository         repository.OrderRepository
}

func (service *trackingOrderServiceImpl) Create(trackingOrderService model.CreateTrackingOrderRequest) error {
	exists := service.OrderRepository.GetById(trackingOrderService.OrderId)
	if exists == (entity.Order{}) {
		return errors.New("order id not! cannot add tracking order")
	}
	data := entity.TrackingOrder{
		Id:          uuid.New().String(),
		OrderId:     trackingOrderService.OrderId,
		CheckPoints: trackingOrderService.CheckPoints,
		TimeStamp:   time.Now(),
	}
	err := service.TrackingOrderRepository.Create(data)
	if err != nil {
		return err
	}

	return nil
}
