package controller

import (
	"luwjistik/config"
	"luwjistik/repository"
	"luwjistik/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func createTestApp() *fiber.App {
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	userController.Route(app)
	orderController.Route(app)
	trackingOrderController.Route(app)
	return app
}

var configuration = config.New("../.env.test")

var database = config.SetUpGormDB(configuration)
var redisClient = config.NewRedisClient(configuration)

var userRepository = repository.NewUserRepository(database)
var orderRepository = repository.NewOrderRepository(database)
var trackingOrderRepository = repository.NewTrackingOrderRepository(database)

var userService = service.NewUserService(&userRepository)
var orderService = service.NewOrderService(&orderRepository)
var trackingOrderService = service.NewTrackingOrderService(&trackingOrderRepository, &orderRepository)

var userController = NewUserController(&userService, redisClient)
var orderController = NewOrderController(&orderService)
var trackingOrderController = NewTrackingOrderController(&trackingOrderService)

var app = createTestApp()
