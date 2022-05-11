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
	"os"
	"strconv"
)

//DONE: add air livereload support
//DONE: add conf system
//DONE: upgrade db
//DONE: add password encryption
//DONE: security update (tls)
//TODO: db migration
//TODO: tests
//TODO: comment and document code
//TODO: JWT validation
//optional
//TODO: upgrade websocket user message gateway
//DONE: add docker support
func main() {

	inFlags := os.Args[1:]

	router := gin.Default()
	router.SetTrustedProxies([]string{})

	//debug config if debug
	var appConf *config.Config
	var err error
	if utils.IfHasFlag(inFlags, "--Prod") || utils.IfHasFlag(inFlags, "--Production") {
		appConf, err = config.Parse("./configProd.yml")
	} else {
		appConf, err = config.Parse("./configDebug.yml")
	}

	if err != nil {
		fmt.Printf("Failed to parse appConf!: %s", err.Error())
		return
	}

	db, err := DB.ConnectDB(appConf.DbAddr)
	//DB.MigrateUsers(db)
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
	router.Use(utils.SetCORS())

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

	if env.Config.UseSSL {
		err = router.RunTLS(":"+strconv.Itoa(appConf.Port), "./SSLCert/https/server.crt", "./SSLCert/https/server.key")
		if err != nil {
			return
		} // listen and serve on 0.0.0.0:81
	} else {
		err = router.Run(":" + strconv.Itoa(appConf.Port))
		if err != nil {
			return
		} // listen and serve on 0.0.0.0:81
	}
}
