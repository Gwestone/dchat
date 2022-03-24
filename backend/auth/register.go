package auth

import (
	"dchat/backend/utils"
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
	_, err = env.Db.Exec("INSERT INTO users (UserId, Username, Password) VALUES (?, ?, ?)", Uuid, registerData.Username, registerData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "username already used",
			"error":   err.Error(),
		})
		return
	}

	//get user to create JWT token
	var user UserSession
	err = env.Db.QueryRow("SELECT * from users WHERE UserId IS ?", Uuid).Scan(&user.Id, &user.UserId, &user.Username, &user.Password)

	token, err := CreateToken(user)
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
