package service

import "luwjistik/model"

type UserService interface {
	Create(user model.CreateUserRequest)
	Login(email, password string) (bool, string)
}
