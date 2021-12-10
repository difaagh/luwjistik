package repository

import (
	"luwjistik/entity"

	"gorm.io/gorm"
)

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{Conn: database}
}

type userRepositoryImpl struct {
	Conn *gorm.DB
}

func (repository *userRepositoryImpl) Create(user entity.User) error {
	if err := repository.Conn.Table(entity.User_table).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repository *userRepositoryImpl) GetByEmail(email string) entity.User {
	var user entity.User
	repository.Conn.Table(entity.User_table).Where("email = ?", email).Find(&user)

	return user
}

func (repository *userRepositoryImpl) DeleteAll() {
	query := "TRUNCATE TABLE " + entity.User_table
	repository.Conn.Exec(query)
}
