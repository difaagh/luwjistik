package util

import (
	"errors"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtWrapper struct {
	SecretKey string
	Issuer    string
}

type JwtClaim struct {
	Email    string
	Name     string
	MobileNo string
	jwt.StandardClaims
}

func (j *JwtWrapper) GenerateToken(name, email, mobileNo string, now time.Time) (signedToken string, err error) {
	claims := &JwtClaim{
		Email:    email,
		Name:     name,
		MobileNo: mobileNo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Local().Add(time.Minute * 5).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {

		return
	}

	return
}

func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		log.Println(err)
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		log.Println(err, "error 2")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		log.Println(err, "error 3")
		return
	}

	return

}
