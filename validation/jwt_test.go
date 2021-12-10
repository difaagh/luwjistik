package validation

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const email = "test@email.com"

var signedToken string

func TestJwt_GenerateJwt(t *testing.T) {
	jwt := JwtWrapper{
		SecretKey: "secret for test",
		Issuer:    "test",
	}
	var err error
	signedToken, err = jwt.GenerateToken(email, time.Now())
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
}
