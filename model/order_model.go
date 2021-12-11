package model

type CreateOrderRequest struct {
	Id               string `json:"id"`
	Weight           uint16 `json:"weight"`
	Sender           string `json:"sender"`
	SenderMobileNo   string `json:"senderMobileNo"`
	ReceiverAddress  string `json:"receiverAddress"`
	ReceiverName     string `json:"receiverName"`
	ReceiverMobileNo string `json:"receiverMobileNo"`
	Status           int16  `json:"status"`
}

type GetOrderDetailRequest struct {
	Id               string                    `json:"id"`
	Weight           uint16                    `json:"weigth"`
	Sender           string                    `json:"sender"`
	SenderMobileNo   string                    `json:"senderMobileNo"`
	ReceiverAddress  string                    `json:"receiverAddress"`
	ReceiverName     string                    `json:"receiverName"`
	ReceiverMobileNo string                    `json:"receiverMobileNo"`
	Status           int16                     `json:"status"`
	TrackingOrders   []GetTrackingOrderRequest `json:"trackingOrders"`
}
