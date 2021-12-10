package entity

const Order_table = "`order`"

type Order struct {
	Id               string `gorm:"id,omitempty,primary_key"`
	Sender           string `gorm:"sender,omitmepty"`
	Weight           uint16 `gorm:"weight,omitempty"`
	SenderMobileNo   string `gorm:"sender_mobile_no"`
	ReceiverAddress  string `gorm:"receiver_address,omityempty"`
	ReceiverName     string `gorm:"receiver_name"`
	ReceiverMobileNo string `gorm:"receiver_mobile_no"`
	Status           string `gorm:"status,omitempty"`
}

type OrderDetail struct {
	Order
	// order id one to many to TrackingOrder struct
	TrackingOrders []TrackingOrder `gorm:"ForeignKey:OrderId"`
}

func (Order) TableName() string {
	return Order_table
}
