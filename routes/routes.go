package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"omochi/app/controller"
	"omochi/app/controller/auth"
	"omochi/config"
	"omochi/middleware"
)

func Router(){
	engine := gin.Default()
	engine.Use(middleware.ErrorMiddleware())
	APIEngine := engine.Group("/v1")
	APIEngine.GET("/health", controller.Health)
	{
		auth.Router(APIEngine)
	}
	engine.Run(fmt.Sprintf(":%d",config.Config.Port))
}
