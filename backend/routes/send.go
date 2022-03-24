package routes

import (
	"dchat/backend/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type dataJSON struct {
	Message string `json:"message"`
}

func Send(c *gin.Context) {

	//get data from request body
	var JSONData dataJSON
	c.BindJSON(&JSONData)

	//get data from params
	receiver, _ := strconv.Atoi(c.Param("receiver"))

	//get global variables
	env := utils.GetEnv(c)

	//get data from token
	rawClaims, exists := c.Get("tokenData")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant get token",
		})
		return
	}
	claims := rawClaims.(jwt.MapClaims)
	sender_, ok := claims["Id"].(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant parse id from token",
		})
		return
	}
	sender := int(sender_)

	_, err := env.Db.Exec("INSERT INTO messages (\"from\", \"to\", message, date) VALUES (?, ?, ?, ?)", sender, receiver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cant get users",
			"error":   err.Error(),
		})
		return
	}

	var messageArr []utils.MessageData

	c.JSON(http.StatusOK, gin.H{
		"message":  "successfully processed",
		"messages": messageArr,
	})
}
