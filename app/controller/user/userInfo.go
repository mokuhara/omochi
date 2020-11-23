package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"omochi/app/models"
	"omochi/app/repository"
	"omochi/middleware"
)

func CreateUserInfo(c *gin.Context) {
	userInfo := models.UserInfo{}
	err := c.BindJSON(&userInfo)
	if err != nil {
		log.Println("action=CreateUserInfo bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	userInfoRepository := repository.UserInfoRepository{}
	if err = userInfoRepository.Create(&userInfo); err != nil {
		log.Println("action=CreateUserInfo failed to create userInfo")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}


func Router(group *gin.RouterGroup) {
	usrInfoEngine := group.Group("/userInfo/:id")
	usrInfoEngine.Use(middleware.IsLogin())
	{
		usrInfoEngine.POST("/create", CreateUserInfo)
	}
}