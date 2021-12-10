package model

type CreateOrderRequest struct {
	Id               string
	Weight           uint16
	Sender           string
	SenderMobileNo   string
	ReceiverAddress  string
	ReceiverName     string
	ReceiverMobileNo string
	Status           string
}

type GetOrderDetailRequest struct {
	Id               string
	Weight           uint16
	Sender           string
	SenderMobileNo   string
	ReceiverAddress  string
	ReceiverName     string
	ReceiverMobileNo string
	Status           string
	TrackingOrders   []GetTrackingOrderRequest
}
