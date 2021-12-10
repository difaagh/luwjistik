package main

import (
	"luwjistik/config"
	"luwjistik/controller"
	"luwjistik/exception"
	"luwjistik/middleware"
	"luwjistik/model"
	"luwjistik/repository"
	"luwjistik/service"
	"luwjistik/validation"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Setup Configuration
	configuration := config.New()
	redis := config.NewRedisClient(configuration)

	database := config.SetUpGormDB(configuration)

	// Setup Repository
	userRepository := repository.NewUserRepository(database)
	orderRepository := repository.NewOrderRepository(database)
	trackingOrderRepository := repository.NewTrackingOrderRepository(database)

	// Setup Service
	userService := service.NewUserService(&userRepository)
	orderService := service.NewOrderService(&orderRepository)
	trackingOrderService := service.NewTrackingOrderService(&trackingOrderRepository, &orderRepository)

	// Setup Controller
	userController := controller.NewUserController(&userService, redis)
	orderController := controller.NewOrderController(&orderService)
	trackingOrderController := controller.NewTrackingOrderController(&trackingOrderService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("ok")
	})

	// Setup Routing
	userController.Route(app)

	app.Use(func(c *fiber.Ctx) error {

		token := c.Get("Authorization")

		err, statusCode := middleware.CheckJwt(token, &validation.JwtWrapper{
			SecretKey: "luwjistik secret",
			Issuer:    "app",
		})
		if err != nil {
			return c.Status(statusCode).JSON(model.WebResponse{
				Code:   statusCode,
				Status: "",
				Data:   err.Error(),
			})
		}
		return c.Next()
	})
	orderController.Route(app)
	trackingOrderController.Route(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	// Start App
	err := app.Listen(":" + port)
	exception.PanicIfNeeded(err)
}
