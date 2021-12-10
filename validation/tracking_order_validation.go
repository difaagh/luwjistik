package validation

import (
	"luwjistik/exception"
	"luwjistik/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateCreateTrackingOrder(request model.CreateTrackingOrderRequest) {

	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required, validation.Length(1, 36)),
		validation.Field(&request.OrderId, validation.Required, validation.Length(1, 36)),
		validation.Field(&request.TimeStamp, validation.Required),
		validation.Field(&request.CheckPoints, validation.Required, validation.Length(1, 250)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
