package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohammedaugi13/simple-rest-go/pkg/app"
	"github.com/mohammedaugi13/simple-rest-go/pkg/lib"
)

func Server() {
	_ = godotenv.Load()
	lib.InitAppConfig()
	lib.GetLogger()

	if lib.AppConfig.App.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}
	r := gin.Default()
	r.Use(lib.GinLogger(), gin.Recovery())

	log.Println("🔰 Application Environment: ", lib.AppConfig.App.Env)
	app.BootstrapApp(r)

	err := r.Run(fmt.Sprintf(`:%d`, lib.AppConfig.App.ServerPort))
	if err != nil {
		fmt.Println(err)
		return
	}
}
