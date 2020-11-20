package routes

import (
	"fmt"
	"github.com/gin-contrib/cors"
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
	engine.Use(cors.New(cors.Config{
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		//TODO AllowOriginsがザルなので絞る
		AllowOrigins: []string{
			"*",
		},
	}))
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
