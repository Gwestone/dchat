package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func keyFunc(token *jwt.Token) (interface{}, error) {
	//Make sure that the token method conform to "SigningMethodHMAC"
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(Secret), nil
}

func LoginRequired(c *gin.Context) {

	if c.GetHeader("Authorization") == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "need authorization"})
	}

	rawToken := strings.Split(c.GetHeader("Authorization"), " ")[1]
	token, err := jwt.Parse(rawToken, keyFunc)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "failed to parse jwt(token may be corrupted or session has expired)"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cant parse token calims",
		})
		return
	}

	c.Set("tokenData", claims)
	c.Next()
}
