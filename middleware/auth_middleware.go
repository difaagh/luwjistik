package middleware

import (
	"errors"
	"luwjistik/util"
	"strings"
)

type CheckJwtReturn struct {
	Err        error
	StatusCode int
	Email      string
	Name       string
	MobileNo   string
}

func CheckJwt(_token string, validation *util.JwtWrapper) *CheckJwtReturn {
	var data CheckJwtReturn
	if _token == "" {
		data = CheckJwtReturn{
			Err:        errors.New("token not provided"),
			StatusCode: 403,
		}
		return &data
	}

	var token string
	splittedToken := strings.Split(_token, "Bearer")
	if len(splittedToken) > 1 {
		_t := strings.Split(splittedToken[1], " ")
		token = _t[1]
	} else {
		token = splittedToken[0]
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
