package routes

import (
	"dchat/backend/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Message(c *gin.Context) {
	//get data from params
	receiver := c.Param("receiver")

	//get global variables
	env := utils.GetEnv(c)

	//get jwt token data
	rawClaims, exists := c.Get("tokenData")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant get token",
		})
		return
	}

	claims := rawClaims.(jwt.MapClaims)
	sender := claims["Username"].(string)

	//get datqa from db

	fmt.Printf(sender + "\n" + receiver)

	rows, err := env.Db.Query("SELECT * FROM main.messages WHERE \"From\" = $2 AND \"To\" = $1", sender, receiver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Cant get users",
			"error":   err.Error(),
		})
		return
	}

	var messageArr []utils.MessageData
	for rows.Next() {
		var Id int
		message := utils.NewMessageData()
		rows.Scan(&Id, &message.From, &message.To, &message.Message, &message.Date)
		messageArr = append(messageArr, *message)
	}

	//responce
	c.JSON(http.StatusOK, gin.H{
		"message":  "successfully processed",
		"messages": messageArr,
	})
}
