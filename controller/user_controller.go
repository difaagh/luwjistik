package controller

import (
	"log"
	"luwjistik/exception"
	"luwjistik/model"
	"luwjistik/service"
	"luwjistik/validation"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController struct {
	UserService service.UserService
	Redis       *redis.Client
}

func NewUserController(userService *service.UserService, redis *redis.Client) UserController {
	return UserController{UserService: *userService, Redis: redis}
}

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
}

func (controller *UserController) Register(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)
	if request.Id == "" {
		request.Id = uuid.New().String()

	}

	controller.UserService.Create(request)
	exception.PanicIfNeeded(err)
	return c.Status(204).JSON(model.WebResponse{
		Code:   204,
		Status: "OK",
		Data:   nil,
	})
}

func (controller *UserController) Login(c *fiber.Ctx) error {
	var request model.LoginUserRequest
	err := c.BodyParser(&request)

	exception.PanicIfNeeded(err)

	if err, responseStr := controller.UserService.Login(request.Email, request.Password); err != true {
		log.Println(responseStr)
		return c.Status(400).JSON(model.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   responseStr,
		})
	}

	jwt := validation.JwtWrapper{
		SecretKey: "luwjistik secret",
		Issuer:    "app",
	}
	var token string
	// get token in redis before generate new token
	redisResult := controller.Redis.Get(c.Context(), request.Email)
	oldToken, _ := redisResult.Result()
	if oldToken == "" {
		newToken, err := jwt.GenerateToken(request.Email, time.Now())
		exception.PanicIfNeeded(err)
		token = newToken
		controller.Redis.Set(c.Context(), request.Email, token, time.Minute*5)
	} else {
		token = oldToken
	}
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   token,
	})
}
