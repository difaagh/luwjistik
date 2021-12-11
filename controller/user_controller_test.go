package controller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"luwjistik/entity"
	"luwjistik/model"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestUserController_Login(t *testing.T) {
	userRepository.DeleteAll()
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

	requestBody, _ := json.Marshal(&loginRequest)

	request := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)
	t.Log(response.StatusCode)

	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	t.Log(webResponse.Data)
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "OK", webResponse.Status)
	assert.NotEqual(t, "", webResponse.Data)

}

func TestUserController_Register(t *testing.T) {
	userRepository.DeleteAll()

	registerRequest := model.CreateUserRequest{
		Id:       uuid.New().String(),
		Name:     "john doe",
		Email:    "john_doe@email.com",
		Password: "dark secret",
	}

	requestBody, _ := json.Marshal(registerRequest)

	request := httptest.NewRequest("POST", "/api/register", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 204, response.StatusCode)

	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
}
