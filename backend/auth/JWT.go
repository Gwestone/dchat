package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const Secret = "123SUPERSECRET123"

func GenToken(user UserSession) (string, error) {

	atClaims := jwt.MapClaims{
		"UserId":   user.UserId,
		"Username": user.Username,
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
