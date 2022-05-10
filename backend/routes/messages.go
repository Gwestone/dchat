package routes

import (
	"github.com/Gwestone/dchat/auth"
	"github.com/Gwestone/dchat/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Messages(c *gin.Context) {
	//get global variables
	env := utils.GetEnv(c)

	//get JWT token data
	rawClaims, exists := c.Get("tokenData")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant get token",
		})
		return
	}
	claims := rawClaims.(jwt.MapClaims)

	userId, ok := claims["UserId"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant parse user id from token",
		})
	}

	//get data from db
	rows, err := env.Db.Query("SELECT * from main.users WHERE \"UserId\" != $1", userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Canrt get users",
		})
		return
	}

	var Password, UserId string
	var userArr []auth.PublicUserData
	for rows.Next() {
		user := auth.NewPublicUserData()
		rows.Scan(&user.Id, &user.Username, &Password, &UserId)
		userArr = append(userArr, *user)
	}

	//responce
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully processed",
		"users":   userArr,
	})
}
