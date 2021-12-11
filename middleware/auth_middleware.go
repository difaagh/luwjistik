package middleware

import (
	"errors"
	"luwjistik/util"
)

type CheckJwtReturn struct {
	Err        error
	StatusCode int
	Email      string
	Name       string
	MobileNo   string
}

func CheckJwt(token string, validation *util.JwtWrapper) *CheckJwtReturn {
	var data CheckJwtReturn
	if token == "" {
		data = CheckJwtReturn{
			Err:        errors.New("token not provided !"),
			StatusCode: 403,
		}
		return &data
	}
	claims, err := validation.ValidateToken(token)
	if err != nil {
		data = CheckJwtReturn{
			Err:        err,
			StatusCode: 400,
		}
		return &data
	}
	data = CheckJwtReturn{
		Email:    claims.Email,
		Name:     claims.Name,
		MobileNo: claims.MobileNo,
	}
	return &data
}
