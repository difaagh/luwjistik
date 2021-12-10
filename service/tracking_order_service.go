package service

import "luwjistik/model"

type TrackingOrderService interface {
	Create(trackingOrder model.CreateTrackingOrderRequest) error
}
