package main

import (
	"context"
	"fmt"
	"github.com/Gwestone/dchat/DB"
	"github.com/Gwestone/dchat/auth"
	"github.com/Gwestone/dchat/config"
	"github.com/Gwestone/dchat/routes"
	"github.com/Gwestone/dchat/session"
	"github.com/Gwestone/dchat/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

//DONE: add air livereload support
//DONE: add conf system
//DONE: upgrade db
//TODO: add password encryption
//TODO: upgrade websocket user message gateway
//TODO: add log system
//TODO: add migration system
//TODO: add docker support
//TODO: security update db(tls)
//TODO: tests
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

	rdb := session.Connect(appConf.RedisAddr, appConf.RedisPassword)
	ctx := context.Background()

	env := utils.NewEnv()
	cache := session.NewStorageContext(rdb, ctx)

	env.Db = db
	env.Config = appConf
	env.SessionCtx = cache

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

	err = router.Run(":" + strconv.Itoa(appConf.Port))
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:81
}
