package auth

import (
	"dchat/backend/utils"
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
	err = env.Db.QueryRow("SELECT * from users WHERE Username=? AND Password=?", loginData.Username, loginData.Password).Scan(&user.Id, &user.UserId, &user.Username, &user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to find user in db",
			"error":   err.Error(),
		})
		return
	}

	//create jwt token and send it to client
	token, err := CreateToken(user)
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
