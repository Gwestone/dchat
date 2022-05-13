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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username already used",
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"message": "request processed successfully",
	})
}
