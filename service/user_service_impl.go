package service

import (
	"log"
	"luwjistik/entity"
	"luwjistik/exception"
	"luwjistik/model"
	"luwjistik/repository"
	"luwjistik/validation"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func (service *userServiceImpl) Create(user model.CreateUserRequest) {
	validation.ValidateCreateUser(user)
	validation.ValidateEmail(user.Email)

	exists := service.UserRepository.GetByEmail(user.Email)
	if exists != (entity.User{}) {
		panic(exception.ValidationError{
			Message: "email already taken",
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
	service.UserRepository.Create(entity.User{
		Id:       uuid.New().String(),
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
	})
}

func (service *userServiceImpl) Login(email, password string) (bool, string) {
	if email == "" || password == "" {
		return false, "password or email cannot be blank"
	}
	user := service.UserRepository.GetByEmail(email)
	if user == (entity.User{}) {
		return false, "password or username invalid"
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err.Error())
		return false, "password or username invalid"
	}
	return true, ""
}
