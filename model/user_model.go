package model

type CreateUserRequest struct {
	Id       string
	Name     string
	Email    string
	Password string
}

type LoginUserRequest struct {
	Email    string
	Password string
}
