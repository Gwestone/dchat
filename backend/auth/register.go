package auth

import (
	"fmt"
	"github.com/Gwestone/dchat/cryptoMod"
	"github.com/Gwestone/dchat/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
)

func Register(c *gin.Context) {

	//get data from req body
	var registerData userDataJSON
	err := c.BindJSON(&registerData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cant parse json",
		})
		return
	}

	err = validator.New().Struct(registerData)
	//TODO specify what is wrong
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input data provided",
		})
		return
	}

	//get global var
	env := utils.GetEnv(c)

	//insert user into db
	hashedPassword := cryptoMod.HashPassword(registerData.Password)
	Uuid := uuid.New().String()

	_, err = env.Db.Exec("INSERT INTO main.users (\"Username\", \"Password\", \"UserId\") VALUES ($1, $2, $3)", registerData.Username, hashedPassword, Uuid)
	if err != nil {
		fmt.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "username already used",
		})
		return
	}

	//get user to create JWT token
	var user UserSession
	err = env.Db.QueryRow("SELECT * from main.users WHERE \"UserId\" == $1", Uuid).Scan(&user.Id, &user.Username, &user.Password, &user.UserId)

	if !EmptyValidate(user) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "cant create user",
		})
		return
	}
	token, err := GenToken(user, env.Config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to generate",
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"message": "request processed successfully",
		"token":   token,
	})
}
