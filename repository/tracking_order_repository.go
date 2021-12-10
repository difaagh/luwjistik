package repository

import "luwjistik/entity"

type TrackingOrderRepository interface {
	Create(trackingOrder entity.TrackingOrder) error

	DeleteAll()
}
