package model

import "time"

type CreateTrackingOrderRequest struct {
	Id          string    `json:"id"`
	OrderId     string    `json:"orderId"`
	CheckPoints string    `json:"checkpoints"`
	TimeStamp   time.Time `json:"timeStamp"`
}

type GetTrackingOrderRequest struct {
	Id          string    `json:"id"`
	OrderId     string    `json:"orderId"`
	CheckPoints string    `json:"checkpoints"`
	TimeStamp   time.Time `json:"timeStamp"`
}
