package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const Secret = "123SUPERSECRET123"

func CreateToken(user UserSession) (string, error) {

	atClaims := jwt.MapClaims{
		"Id":       user.Id,
		"UserId":   user.UserId,
		"Username": user.Username,
		"Password": user.Password,
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
