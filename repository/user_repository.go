package repository

import "luwjistik/entity"

type UserRepository interface {
	Create(user entity.User) error

	GetByEmail(email string) (user entity.User)

	DeleteAll()
}
