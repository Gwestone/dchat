package main

import (
	"dchat/backend/DB"
	_ "dchat/backend/DB"
	"dchat/backend/auth"
	"dchat/backend/routes"
	"dchat/backend/utils"
	"github.com/gin-gonic/gin"
)

//TODO: upgrade session management system to localCached
func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{})

	db := DB.ConnectDB()
	env := utils.NewEnv()
	env.Db = db

	//DB.MigrateUsers(db)

	router.Use(env.SetEnv)
	authRoute := router.Group("/auth")
	{
		authRoute.POST("/login", auth.Login)
		authRoute.POST("/register", auth.Register)
	}

	messageRoute := router.Group("/messages")
	{
		messageRoute.POST("/all", auth.LoginRequired, routes.Messages)
		messageRoute.POST("/:receiver", auth.LoginRequired, routes.Message)
		//messageRoute.POST("/send/:receiver", auth.LoginRequired, routes.Send)
	}

	err := router.Run(":81")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:81
}
