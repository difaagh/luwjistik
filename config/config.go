package config

import (
	"luwjistik/exception"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	production := os.Getenv("GO_ENV")
	if production != "production" {
		err := godotenv.Load(filenames...)
		exception.PanicIfNeeded(err)
	}
	return &configImpl{}
}
