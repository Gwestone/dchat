package routes

import (
	"fmt"
	"github.com/Gwestone/dchat/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type dataJSON struct {
	Message string `json:"message"`
}

func Send(c *gin.Context) {

	//get data from request body
	var JSONData dataJSON
	err := c.BindJSON(&JSONData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to parse body",
		})
		return
	}

	//get data from params
	receiver := c.Param("receiver")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant parse user id",
		})
		return
	}

	//get global variables
	env := utils.GetEnv(c)

	//get data from JWT token
	rawClaims, exists := c.Get("tokenData")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant get token",
		})
		return
	}
	claims := rawClaims.(jwt.MapClaims)
	sender := claims["Username"].(string)

	fmt.Printf(receiver + "\n")
	fmt.Printf(sender)

	//insert data into db
	_, err = env.Db.Exec("INSERT INTO main.messages (\"From\", \"To\", \"Message\", \"Date\") VALUES ($1, $2, $3, $4)", sender, receiver, JSONData.Message, time.Now().Unix())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to insert data",
		})
		return
	}

	//responce
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully processed",
	})
}
