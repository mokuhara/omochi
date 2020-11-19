package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"omochi/app/controller"
	"omochi/app/controller/auth"
	"omochi/app/repository"
	"omochi/config"
	"omochi/middleware"
	"omochi/utils"
)

func main(){
	utils.LoggingSettings(config.Config.LogFile)
	repository.DBCon()

	engine := gin.Default()
	engine.Use(middleware.ErrorMiddleware())
	APIEngine := engine.Group("/v1")
	APIEngine.GET("/health", controller.Health)
	{
		authEngine := APIEngine.Group("/auth")
		{
			authEngine.POST("/signup", auth.Signup)
			authEngine.POST("/login", auth.Login)
		}
	}
	engine.Run(fmt.Sprintf(":%d",config.Config.Port))
}