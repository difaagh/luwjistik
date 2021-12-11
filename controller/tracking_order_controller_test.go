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

func TestTrackingOrder_CreateTrackingOrder(t *testing.T) {
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

	order := entity.Order{
		Id:               uuid.New().String(),
		Weight:           5,
		Sender:           user.Name,
		SenderEmail:      user.Email,
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
	request.Header.Set("Authorization", webResponseLogin.Data.(string))

	response, _ := app.Test(request)
	assert.Equal(t, 204, response.StatusCode)

}
