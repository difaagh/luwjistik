package controller

import (
	"context"
	"log"
	"luwjistik/config"
	"luwjistik/middleware"
	"luwjistik/model"
	"luwjistik/repository"
	"luwjistik/service"
	"luwjistik/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func createTestApp() *fiber.App {
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
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
		log.Println(dataUser)
		ctx := context.WithValue(c.Context(), util.UserContextValue, dataUser)
		c.SetUserContext(ctx)
		return c.Next()
	})
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
