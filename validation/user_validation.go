package validation

import (
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
		validation.Field(&request.MobileNo, validation.Length(0, 20)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateEmail(email string) {

	_, err := mail.ParseAddress(email)
	if err != nil {
		panic(exception.ValidationError{
			Message: "Invalid email",
		})
	}
}
