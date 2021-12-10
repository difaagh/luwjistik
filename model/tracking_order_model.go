package model

import "time"

type CreateTrackingOrderRequest struct {
	Id          string
	OrderId     string
	CheckPoints string
	TimeStamp   time.Time
}

type GetTrackingOrderRequest struct {
	Id          string
	OrderId     string
	CheckPoints string
	TimeStamp   time.Time
}
