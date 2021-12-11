package controller

import (
	"luwjistik/exception"
	"luwjistik/model"
	"luwjistik/service"
	"luwjistik/util"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrderController struct {
	OrderService service.OrderService
}

func NewOrderController(orderService *service.OrderService) OrderController {
	return OrderController{OrderService: *orderService}
}

func (controller *OrderController) Route(app *fiber.App) {
	app.Post("/api/order", controller.CreateOrder)
	app.Get("/api/order/:id", controller.GetDetailOrder)
}

func (controller *OrderController) CreateOrder(c *fiber.Ctx) error {
	var request model.CreateOrderRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	if request.Id == "" {
		request.Id = uuid.New().String()

	}
	if request.Status == 0 {
		request.Status = 1
	}

	email := c.UserContext().Value(util.UserContextValue).(util.ContextValues)[util.UserEmailValue].(string)
	name := c.UserContext().Value(util.UserContextValue).(util.ContextValues)[util.UserNameValue].(string)
	mobileNo := c.UserContext().Value(util.UserContextValue).(util.ContextValues)[util.UserMobileNoValue].(string)
	request.Sender = name
	request.SenderEmail = email
	request.SenderMobileNo = mobileNo

	controller.OrderService.Create(request)
	message := struct {
		OrderId string `json:"orderId"`
	}{
		OrderId: request.Id,
	}
	return c.Status(200).JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   message,
	})
}

func (controller *OrderController) GetDetailOrder(c *fiber.Ctx) error {
	orderId := c.Params("id")

	email := c.UserContext().Value(util.UserContextValue).(util.ContextValues)[util.UserEmailValue].(string)
	response := controller.OrderService.GetDetailById(orderId, email)

	return c.Status(200).JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}
