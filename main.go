package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tiketin-management-api-with-go/config"
	"github.com/tiketin-management-api-with-go/database"
	"github.com/tiketin-management-api-with-go/routes"
)

func main() {
	config.Init()
	database.Init()

	web := gin.Default()

	routes.InitRoute(web)

	web.Run(viper.GetString("App.Port"))
}
