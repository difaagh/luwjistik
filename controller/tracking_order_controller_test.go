package controller

import (
	"bytes"
	"encoding/json"
	"luwjistik/entity"
	"luwjistik/model"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTrackingOrder_CreateTrackingOrder(t *testing.T) {
	orderRepository.DeleteAll()
	trackingOrderRepository.DeleteAll()

	order := entity.Order{
		Id:               uuid.New().String(),
		Weight:           5,
		Sender:           "John Doe",
		SenderMobileNo:   "6281222",
		ReceiverAddress:  "Jakarta, Indonesia",
		ReceiverName:     "Keith Heart",
		ReceiverMobileNo: "681111",
	}

	orderRepository.Create(entity.OrderDetail{Order: order})

	trackingOrder := model.CreateTrackingOrderRequest{
		Id:          uuid.New().String(),
		OrderId:     order.Id,
		CheckPoints: "Order Created",
		TimeStamp:   time.Now(),
	}

	requestBody, _ := json.Marshal(&trackingOrder)

	request := httptest.NewRequest("POST", "/api/order/tracking", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 204, response.StatusCode)

}
