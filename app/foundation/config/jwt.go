package config

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var SigningKey []byte

func CreateJWTToken(Issuer string) (*jwt.Token, string, error) {
	expiresAt := time.Now().Add(time.Hour * 24).Unix()
	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: expiresAt,
		Issuer:    Issuer,
		Id:        uuid.NewString(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SigningKey)
	return token, ss, err
}
