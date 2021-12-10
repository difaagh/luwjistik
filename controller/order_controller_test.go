package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"luwjistik/entity"
	"luwjistik/model"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestOrderController_CreateOrder(t *testing.T) {
	orderRepository.DeleteAll()

	orderRequest := model.CreateOrderRequest{
		Weight:           5,
		Sender:           "John Doe",
		SenderMobileNo:   "6281222",
		ReceiverAddress:  "Jakarta, Indonesia",
		ReceiverName:     "Keith Heart",
		ReceiverMobileNo: "681111",
		Status:           1,
	}

	requestBody, _ := json.Marshal(orderRequest)

	request := httptest.NewRequest("POST", "/api/order", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)
	t.Log(response.Status)

	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)

	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	obj := webResponse.Data.(interface{})
	jsonData, _ := json.Marshal(obj)
	type data struct {
		OrderId string
	}
	orderObj := data{}
	json.Unmarshal(jsonData, &orderObj)
	assert.NotEqual(t, "", orderObj.OrderId)
}

func TestOrderController_GetDetailOrder(t *testing.T) {
	orderRepository.DeleteAll()
	trackingOrderRepository.DeleteAll()

	orderId := uuid.New().String()
	order := entity.Order{
		Id:               orderId,
		Weight:           5,
		Sender:           "John Doe",
		SenderMobileNo:   "6281222",
		ReceiverAddress:  "Jakarta, Indonesia",
		ReceiverName:     "Keith Heart",
		ReceiverMobileNo: "681111",
	}
	trackingOrder := entity.TrackingOrder{
		Id:          uuid.New().String(),
		OrderId:     orderId,
		CheckPoints: "Order Created",
		TimeStamp:   time.Now(),
	}
	orderRepository.Create(entity.OrderDetail{Order: order})
	trackingOrderRepository.Create(trackingOrder)

	request := httptest.NewRequest("GET", "/api/order/"+orderId, nil)
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)

	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)

	obj := webResponse.Data.(interface{})
	jsonData, _ := json.Marshal(obj)
	orderObj := model.GetOrderDetailRequest{}
	json.Unmarshal(jsonData, &orderObj)
	assert.NotEqual(t, "", orderObj.Id)

	list := orderObj.TrackingOrders
	containstTrackingOrder := false

	for _, trackingOrder := range list {

		if trackingOrder.Id != "" {
			containstTrackingOrder = true
		}
	}
	assert.True(t, containstTrackingOrder)
}
