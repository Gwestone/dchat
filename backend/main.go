package main

import (
	"dchat/backend/DB"
	"dchat/backend/auth"
	"dchat/backend/config"
	"dchat/backend/routes"
	"dchat/backend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

//TODO: upgrade session management system to localCached(redis)
//TODO: upgrade user message gateway
//DONE: upgrade db
//TODO: add log system
//TODO: add migration system
//TODO: add docker support
//TODO: add air support
//DONE: add conf system
func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{})

	appConf, err := config.Parse("./config.yml")
	if err != nil {
		fmt.Printf("Failed to parse appConf!: %s", err.Error())
		return
	}

	db, err := DB.ConnectDB(appConf.DbAddr)
	if err != nil {
		fmt.Printf("Failed to connect db: %s", err.Error())
		return
	}

	env := utils.NewEnv()
	env.Db = db
	env.Config = appConf

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
		messageRoute.POST("/send/:receiver", auth.LoginRequired, routes.Send)
	}

	//cryptRoute := router.Group("/crypt")
	//{
	//	cryptRoute.GET("/pubKey", crypto.GetPubKey)
	//}

	err = router.Run(":" + strconv.Itoa(appConf.Port))
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:81
}
