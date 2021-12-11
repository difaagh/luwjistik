package main

import (
	"context"
	"luwjistik/config"
	"luwjistik/controller"
	"luwjistik/exception"
	"luwjistik/middleware"
	"luwjistik/model"
	"luwjistik/repository"
	"luwjistik/service"
	"luwjistik/util"
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
		claims := middleware.CheckJwt(token, &util.JwtWrapper{
			SecretKey: "luwjistik secret",
			Issuer:    "app",
		})

		if claims.Err != nil {
			return c.Status(claims.StatusCode).JSON(model.WebResponse{
				Code:   claims.StatusCode,
				Status: "",
				Data:   claims.Err.Error(),
			})
		}
		dataUser := util.ContextValues{
			util.UserEmailValue:    claims.Email,
			util.UserNameValue:     claims.Name,
			util.UserMobileNoValue: claims.MobileNo,
		}

		ctx := context.WithValue(c.Context(), util.UserContextValue, dataUser)
		c.SetUserContext(ctx)
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
