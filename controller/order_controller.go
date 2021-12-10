package controller

import (
	"luwjistik/exception"
	"luwjistik/model"
	"luwjistik/service"

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

	response := controller.OrderService.GetDetailById(orderId)

	return c.Status(200).JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})

}