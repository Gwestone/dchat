package routes

import (
	"dchat/backend/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Message(c *gin.Context) {
	receiver, _ := strconv.Atoi(c.Param("receiver"))

	env := utils.GetEnv(c)

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

	rows, err := env.Db.Query("SELECT * FROM messages WHERE \"from\" IS ? AND \"to\" IS ?", sender, receiver)
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

	c.JSON(http.StatusOK, gin.H{
		"message":  "successfully processed",
		"messages": messageArr,
	})
}
