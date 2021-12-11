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
	"golang.org/x/crypto/bcrypt"
)

func TestOrderController_CreateOrder(t *testing.T) {
	userRepository.DeleteAll()
	orderRepository.DeleteAll()

	password := "dark secret"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := entity.User{
		Id:       uuid.New().String(),
		Name:     "John Doe",
		Email:    "john_doe@email.com",
		Password: string(hashedPassword),
	}
	userRepository.Create(user)

	orderRequest := model.CreateOrderRequest{
		Weight:           5,
		ReceiverAddress:  "Jakarta, Indonesia",
		ReceiverName:     "Keith Heart",
		ReceiverMobileNo: "681111",
		Status:           1,
	}

	loginRequest := model.LoginUserRequest{
		Password: password,
		Email:    user.Email,
	}

	requestLoginBody, _ := json.Marshal(loginRequest)
	requestLogin := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(requestLoginBody))
	requestLogin.Header.Set("Content-Type", "application/json")
	requestLogin.Header.Set("Accept", "application/json")

	responseLogin, _ := app.Test(requestLogin)

	responseBodyLogin, _ := ioutil.ReadAll(responseLogin.Body)

	webResponseLogin := model.WebResponse{}
	json.Unmarshal(responseBodyLogin, &webResponseLogin)

	requestBody, _ := json.Marshal(orderRequest)

	request := httptest.NewRequest("POST", "/api/order", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", webResponseLogin.Data.(string))

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)

	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	t.Log(webResponse)

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
	userRepository.DeleteAll()
	orderRepository.DeleteAll()
	trackingOrderRepository.DeleteAll()

	password := "dark secret"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := entity.User{
		Id:       uuid.New().String(),
		Name:     "John Doe",
		Email:    "john_doe@email.com",
		Password: string(hashedPassword),
	}
	userRepository.Create(user)

	loginRequest := model.LoginUserRequest{
		Password: password,
		Email:    user.Email,
	}

	requestLoginBody, _ := json.Marshal(loginRequest)
	requestLogin := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(requestLoginBody))
	requestLogin.Header.Set("Content-Type", "application/json")
	requestLogin.Header.Set("Accept", "application/json")

	responseLogin, _ := app.Test(requestLogin)

	responseBodyLogin, _ := ioutil.ReadAll(responseLogin.Body)

	webResponseLogin := model.WebResponse{}
	json.Unmarshal(responseBodyLogin, &webResponseLogin)

	orderId := uuid.New().String()
	order := entity.Order{
		Id:               orderId,
		Weight:           5,
		Sender:           user.Name,
		SenderEmail:      user.Email,
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
	request.Header.Set("Authorization", webResponseLogin.Data.(string))

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)

	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	t.Log(webResponseLogin.Data)
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
