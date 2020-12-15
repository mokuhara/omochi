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

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // MEMO: 本番はOriginが異なるので環境変数で対応する？
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	APIEngine := engine.Group("/v1")
	APIEngine.GET("/health", controller.Health)
	{
		auth.Router(APIEngine)
		user.Router(APIEngine)
		specialist.Router(APIEngine)
		mypage.ProfileRouter(APIEngine)
		mypage.BizPackRouter(APIEngine)
		mypage.PortfolioRouter(APIEngine)
		admin.Router(APIEngine)
	}

	engine.Run(fmt.Sprintf(":%d",config.Config.Port))
}
