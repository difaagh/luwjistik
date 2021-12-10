package controller

import (
	"luwjistik/exception"
	"luwjistik/model"
	"luwjistik/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TrackingOrderController struct {
	TrackingOrderService service.TrackingOrderService
}

func NewTrackingOrderController(trackingOrderService *service.TrackingOrderService) TrackingOrderController {
	return TrackingOrderController{TrackingOrderService: *trackingOrderService}
}

func (controller *TrackingOrderController) Route(app *fiber.App) {
	app.Post("/api/order/tracking", controller.CreateTrackingOrder)
}

func (controller *TrackingOrderController) CreateTrackingOrder(c *fiber.Ctx) error {
	var request model.CreateTrackingOrderRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	if request.Id == "" {
		request.Id = uuid.New().String()

	}
	if request.TimeStamp.IsZero() {
		request.TimeStamp = time.Now()
	}

	err = controller.TrackingOrderService.Create(request)
	if err != nil {
		return c.Status(400).JSON(model.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	return c.Status(204).JSON(model.WebResponse{})
}
