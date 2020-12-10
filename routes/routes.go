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
	"omochi/app/controller/user"
	"omochi/config"
	"omochi/middleware"
)

func Router(){
	engine := gin.Default()
	engine.Use(middleware.ErrorMiddleware())
	//engine.Use(cors.Default())
	engine.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Origin",
			"Access-Control-Request-Method",
			"Access-Control-Request-Headers",
			"Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"authorization",
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
		user.Router(APIEngine)
		specialist.Router(APIEngine)
		mypage.ProfileRouter(APIEngine)
		mypage.BizPackRouter(APIEngine)
		admin.Router(APIEngine)
	}
	engine.Run(fmt.Sprintf(":%d",config.Config.Port))
}
