package auth

import (
	"fmt"
	"github.com/Gwestone/dchat/utils"
	"github.com/gin-gonic/gin"
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
			"error":   err.Error(),
		})
		return
	}

	//get global var
	env := utils.GetEnv(c)

	//insert user into db
	Uuid := uuid.New().String()
	_, err = env.Db.Exec("INSERT INTO main.users (\"Username\", \"Password\", \"UserId\") VALUES ($1, $2, $3)", registerData.Username, registerData.Password, Uuid)
	if err != nil {
		fmt.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "username already used",
			"error":   err.Error(),
		})
		return
	}

	//get user to create JWT token
	var user UserSession
	err = env.Db.QueryRow("SELECT * from main.users WHERE \"UserId\" IS $1", Uuid).Scan(&user.Id, &user.Username, &user.Password, &user.UserId)

	token, err := GenToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to generate",
			"error":   err.Error(),
		})
		return
	}

	//response
	c.JSON(http.StatusOK, gin.H{
		"message": "request processed successfully",
		"token":   token,
	})
}
