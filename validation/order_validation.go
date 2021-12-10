package validation

import (
	"luwjistik/exception"
	"luwjistik/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateCreateOrder(request model.CreateOrderRequest) {

	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required, validation.Length(1, 36)),
		validation.Field(&request.Weight, validation.Required),
		validation.Field(&request.Sender, validation.Required, validation.Length(1, 250)),
		validation.Field(&request.SenderMobileNo, validation.Required, validation.Length(1, 20)),
		validation.Field(&request.ReceiverAddress, validation.Required, validation.Length(1, 250)),
		validation.Field(&request.ReceiverName, validation.Required, validation.Length(1, 250)),
		validation.Field(&request.ReceiverMobileNo, validation.Required, validation.Length(1, 20)),
		validation.Field(&request.Status, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
