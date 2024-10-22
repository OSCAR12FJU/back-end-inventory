package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenUtil struct {
	SecretKey []byte
}

func NewTokenUtil(secretKey string) *TokenUtil {
	return &TokenUtil{SecretKey: []byte(secretKey)}
}

func (t *TokenUtil) GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(t.SecretKey)
}
