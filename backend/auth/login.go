package auth

import (
	"github.com/Gwestone/dchat/cryptoMod"
	"github.com/Gwestone/dchat/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Login(c *gin.Context) {
	//get data from req body
	var loginData userDataJSON
	err := c.BindJSON(&loginData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid json provided",
		})
		return
	}

	err = validator.New().Struct(loginData)
	//TODO specify what is wrong
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input data provided",
		})
		return
	}

	//get global variable
	env := utils.GetEnv(c)

	//get data from db
	var user UserSession
	err = env.Db.QueryRow("SELECT * from main.users WHERE \"Username\"=$1", loginData.Username).Scan(&user.Id, &user.Username, &user.Password, &user.UserId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unable to find user in db",
		})
		return
	}

	err = validator.New().Struct(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "wrong password or username",
		})
		return
	}

	if !cryptoMod.ComparePasswords(user.Password, loginData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "wrong password or username",
		})
		return
	}

	//create jwt token and send it to client
	token, err := GenToken(user, env.Config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "unable to generate",
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
