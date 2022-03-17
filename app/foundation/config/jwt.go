package config

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var SigningKey []byte

func CreateJWTToken(Issuer string) (*jwt.Token, string, error) {
	expiresAt := time.Now().Add(time.Hour * 24)
	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		Issuer:    Issuer,
		ID:        uuid.NewString(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SigningKey)
	return token, ss, err
}
