package repository

import (
	"luwjistik/entity"

	"gorm.io/gorm"
)

func NewOrderRepository(database *gorm.DB) OrderRepository {
	return &orderRepositoryImpl{Conn: database}
}

type orderRepositoryImpl struct {
	Conn *gorm.DB
}

func (repository *orderRepositoryImpl) Create(order entity.OrderDetail) error {
	if err := repository.Conn.Table(entity.Order_table).Create(&order).Error; err != nil {
		return err
	}
	return nil
}

func (repository *orderRepositoryImpl) GetById(orderId string) (order entity.Order) {
	repository.Conn.Where("id = ?", orderId).Find(&order)
	return order
}

func (repository *orderRepositoryImpl) GetDetailById(orderId string) (orders entity.OrderDetail) {
	repository.Conn.Where("id = ?", orderId).Find(&orders)
	repository.Conn.Model(&orders).Association("TrackingOrders").Find(&orders.TrackingOrders)
	return orders
}

func (repository *orderRepositoryImpl) DeleteAll() {
	check := "0"
	querySetCheckFK := "SET FOREIGN_KEY_CHECKS = " + check
	repository.Conn.Exec(querySetCheckFK)
	queryTruncate := "TRUNCATE TABLE " + entity.Order_table + ";"
	repository.Conn.Exec(queryTruncate)
	check = "1"
	querySetCheckFK = "SET FOREIGN_KEY_CHECKS = " + check
	repository.Conn.Exec(querySetCheckFK)
}
