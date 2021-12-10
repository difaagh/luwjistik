package entity

import (
	"time"
)

const Tracking_order_table = "tracking_order"

type TrackingOrder struct {
	Id          string    `gorm:"id,omitempty,primary_key"`
	OrderId     string    `gorm:"order_id,omitempty"`
	CheckPoints string    `gorm:"checkpoints,omitempty"`
	TimeStamp   time.Time `gorm:"time_stamp,omitempty"`
}

func (TrackingOrder) TableName() string {
	return Tracking_order_table
}
