package middleware

import (
	"errors"
	"luwjistik/validation"
)

func CheckJwt(token string, validation *validation.JwtWrapper) (error, int) {
	if token == "" {
		return errors.New("Forbidden"), 403
	}
	_, err := validation.ValidateToken(token)
	if err != nil {
		return err, 400
	}
	return nil, 200
}
