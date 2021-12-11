package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const email = "test@email.com"
const name = "John Doe"
const mobileNo = "0000"

var signedToken string

func TestJwt_GenerateJwt(t *testing.T) {
	jwt := JwtWrapper{
		SecretKey: "secret for test",
		Issuer:    "test",
	}
	var err error
	signedToken, err = jwt.GenerateToken(name, email, mobileNo, time.Now())
	assert.NoError(t, err)
	assert.NotEqual(t, "", signedToken)
}

func TestJwt_ValidateJwt(t *testing.T) {
	jwt := JwtWrapper{
		SecretKey: "secret for test",
		Issuer:    "test",
	}
	claims, err := jwt.ValidateToken(signedToken)
	assert.NoError(t, err)
	assert.Equal(t, email, claims.Email)
	assert.Equal(t, name, claims.Name)
	assert.Equal(t, mobileNo, claims.MobileNo)
}
