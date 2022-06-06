package main

import (
	"context"
	"fmt"
	"github.com/Gwestone/dchat/DB"
	"github.com/Gwestone/dchat/auth"
	"github.com/Gwestone/dchat/color"
	"github.com/Gwestone/dchat/config"
	"github.com/Gwestone/dchat/routes"
	"github.com/Gwestone/dchat/session"
	"github.com/Gwestone/dchat/utils"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
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

	//debug config if debug
	var appConf *config.Config
	var err error

	if utils.IfHasFlag(inFlags, "--Prod") || utils.IfHasFlag(inFlags, "--Production") {
		appConf, err = config.Parse("./configProd.yml")
		gin.SetMode(gin.ReleaseMode)
	} else {
		appConf, err = config.Parse("./configDebug.yml")
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()
	router.SetTrustedProxies([]string{})
	if err != nil {
		fmt.Printf("Failed to parse appConf!: %s", err.Error())
		return
	}

	//error logging to file
	if appConf.UseFileLog {
		abs, err := filepath.Abs("./Log")
		if err != nil {
			err.Error()
		}

		if _, err := os.Stat(abs); !os.IsNotExist(err) {

			now := time.Now()
			milisec := strconv.Itoa(now.Nanosecond() / 1000000)
			filename := now.Format("2006-Jan-Mon-15-04-05") + "-" + milisec + ".log"
			filenameAbs := path.Join(abs, filename)
			f, err := os.OpenFile(filenameAbs, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			gin.DefaultErrorWriter = io.MultiWriter(f)

		} else {
			fmt.Println("cant find path to log folder " + abs)
		}
	}

	db, err := DB.ConnectDB(appConf.DbAddr)
	//DB.MigrateUsers(db)
	if err != nil {
		fmt.Printf("Failed to connect db: %s", err.Error())
		return
	}

	env := utils.NewEnv()

	if appConf.UseRedis {
		fmt.Println("Trying to launch redis service...")
		rdb := session.Connect(appConf.RedisAddr, appConf.RedisPassword)
		ctx := context.Background()
		cache := session.NewCacheContext(rdb, ctx)
		env.SessionCtx = cache
		if rdb.Ping(ctx).Val() == "PONG" {
			fmt.Println("Redis is", color.Green, "online", color.Reset)
		} else {
			fmt.Println("Failed to setup redis on addr :", appConf.RedisAddr)
			return
		}
		fmt.Println("Redis is", color.Green, "online", color.Reset)
	} else {
		fmt.Println("Redis is", color.Red, "offline", color.Reset, "this can cause instability in work or decreased performance")
	}

	env.Db = db
	env.Config = appConf

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
