package config

import (
	"luwjistik/exception"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

func SetUpGormDB(cfg Config) *gorm.DB {

	db, err := NewDb(cfg)
	exception.PanicIfNeeded(err)
	return db
}

func NewDb(c Config) (*gorm.DB, error) {
	return newMysql(c)
}

func newMysql(c Config) (*gorm.DB, error) {
	user := c.Get("MYSQL_USERNAME")
	password := c.Get("MYSQL_PASSWORD")
	host := c.Get("MYSQL_HOST")
	port := c.Get("MYSQL_PORT")
	database := c.Get("MYSQL_DATABASE")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	exception.PanicIfNeeded(err)
	return db, nil
}
