package entity

const User_table = "user_table"

type User struct {
	Id       string `gorm:"id,omitempty,primary_key"`
	Name     string `gorm:"name,omitempty"`
	Password string `gorm:"password,omitempty"`
	Email    string `gorm:"email,omitempty"`
}

func (User) TableName() string {
	return User_table
}
