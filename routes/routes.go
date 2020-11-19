package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"omochi/app/controller"
	"omochi/app/controller/admin"
	"omochi/app/controller/auth"
	"omochi/app/controller/mypage"
	"omochi/app/controller/specialist"
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
		specialist.Router(APIEngine)
		mypage.Router(APIEngine)
		admin.Router(APIEngine)
	}
	engine.Run(fmt.Sprintf(":%d",config.Config.Port))
}
