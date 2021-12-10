package validation

import (
	"log"
	"luwjistik/exception"
	"luwjistik/model"
	"net/mail"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateCreateUser(request model.CreateUserRequest) {

	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Name, validation.Required, validation.Length(1, 250)),
		validation.Field(&request.Email, validation.Required, validation.Length(1, 50)),
		validation.Field(&request.Password, validation.Required, validation.Length(1, 250)),
	)

	if err != nil {
		log.Println("here1")
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateEmail(email string) {

	_, err := mail.ParseAddress(email)
	if err != nil {
		log.Println("here2")
		panic(exception.ValidationError{
			Message: "Invalid email",
		})
	}
}
