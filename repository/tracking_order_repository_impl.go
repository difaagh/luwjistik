package repository

import (
	"luwjistik/entity"

	"gorm.io/gorm"
)

func NewTrackingOrderRepository(database *gorm.DB) TrackingOrderRepository {
	return &trackingOrderRepositoryImpl{Conn: database}
}

type trackingOrderRepositoryImpl struct {
	Conn *gorm.DB
}

func (repository *trackingOrderRepositoryImpl) Create(trackingOrder entity.TrackingOrder) error {
	if err := repository.Conn.Table(entity.Tracking_order_table).Create(&trackingOrder).Error; err != nil {
		return err
	}
	return nil
}

func (repository *trackingOrderRepositoryImpl) DeleteAll() {
	query := "TRUNCATE TABLE " + entity.Tracking_order_table
	repository.Conn.Exec(query)
}
