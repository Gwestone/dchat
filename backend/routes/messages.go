package routes

import (
	"dchat/backend/auth"
	"dchat/backend/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Messages(c *gin.Context) {
	var Password, UserId string
	env := utils.GetEnv(c)

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

	rows, err := env.Db.Query("SELECT * from users WHERE UserId IS NOT ?", userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Canrt get users",
			"error":   err.Error(),
		})
		return
	}

	var userArr []auth.PublicUserData
	for rows.Next() {
		user := auth.NewPublicUserData()
		rows.Scan(&user.Id, &UserId, &user.Username, &Password)
		userArr = append(userArr, *user)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully processed",
		"users":   userArr,
	})
}
