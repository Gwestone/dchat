package auth

import (
	"github.com/Gwestone/dchat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	//get data from req body
	var loginData userDataJSON
	err := c.BindJSON(&loginData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid json provided",
			"error":   err.Error(),
		})
		return
	}

	//get global variable
	env := utils.GetEnv(c)

	//get data from db
	var user UserSession
	err = env.Db.QueryRow("SELECT * from main.users WHERE \"Username\"=$1 AND \"Password\"=$2", loginData.Username, loginData.Password).Scan(&user.Id, &user.Username, &user.Password, &user.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to find user in db",
			"error":   err.Error(),
		})
		return
	}

	//create jwt token and send it to client
	token, err := GenToken(user, env.Config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to generate",
			"error":   err.Error(),
		})
		return
	}

	//responce
	c.JSON(http.StatusOK, gin.H{
		"message": "request processed successfully",
		"token":   token,
	},
	)
}
