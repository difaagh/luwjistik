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
		MobileNo: user.MobileNo,
	})
}

func (service *userServiceImpl) Login(email, password string) (warning string, user model.GetUserRequest) {
	if email == "" || password == "" {
		warning = "password or email cannot be blank"
		return warning, user
	}
	exists := service.UserRepository.GetByEmail(email)
	if exists == (entity.User{}) {
		warning = "password or username invalid"
		return warning, user
	}

	err := bcrypt.CompareHashAndPassword([]byte(exists.Password), []byte(password))
	if err != nil {
		log.Println(err.Error())
		warning = "password or username invalid"
		return warning, user
	}
	user = model.GetUserRequest{
		Id:       exists.Id,
		Name:     exists.Name,
		Email:    exists.Email,
		MobileNo: exists.MobileNo,
	}

	return "", user
}
