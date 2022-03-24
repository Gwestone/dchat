package routes

import (
	"dchat/backend/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
			"error":   err.Error(),
		})
		return
	}

	//get data from params
	receiver, err := strconv.Atoi(c.Param("receiver"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant parse user id",
			"error":   err.Error(),
		})
		return
	}

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

	//insert data into db
	_, err = env.Db.Exec("INSERT INTO messages (\"from\", \"to\", message, date) VALUES (?, ?, ?, ?)", sender, receiver, JSONData.Message, time.Now().Unix())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to insert data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully processed",
	})
}
