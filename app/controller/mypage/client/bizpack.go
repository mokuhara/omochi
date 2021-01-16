package client

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"omochi/app/repository"
	"omochi/app/service"
	"omochi/middleware"
)

var bizpackRepository = repository.BizpackRepository{}

func authenticate(c *gin.Context) (auth *service.Auth, err error) {
	tokenService := service.TokenService{}
	user, err := tokenService.Verify(c)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func Index(c *gin.Context) {
	user, err := authenticate(c)
	println(user)
	if err != nil {
		log.Println("action=DeleteBizpack user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)

		return
	}
	//TODO transaction化していないbizpackのみ取得するよう変更
	bizpacks, err := bizpackRepository.GetAll()

	if err != nil {
		log.Println("action=GetUserBizpacks failed to get bizpacks")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *bizpacks,
	})
}


func ClientBizPackRouter(group *gin.RouterGroup){
	myPageEngine := group.Group("/mypage")
	myPageEngine.Use(middleware.IsLogin())
	{
		ClientEngine := myPageEngine.Group("/client")
		{
			BizpackEngine := ClientEngine.Group("/bizpacks")
			{
				BizpackEngine.GET("", Index)
			}
		}
	}
}
