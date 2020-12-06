package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"omochi/app/repository"
	"omochi/app/service"
)

func IsLogin() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenService := service.TokenService{}
		_, err := tokenService.Verify(c)
		if err != nil {
			log.Println("invalid token")
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"data": err.Error(),
			})
			c.Abort()
		}
	}
}

func IsExistsUserInfo() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenService := service.TokenService{}
		res, _ := tokenService.Verify(c)
		userInfoRepository := repository.UserInfoRepository{}
		userInfo, err := userInfoRepository.GetByUserId(res.UserId)
		if err != nil {
			log.Println("invalid userId")
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"data": err.Error(),
			})
			c.Abort()
		}
		if len(userInfo.Name) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"data": "userInfo is not exist",
			})
			c.Abort()
		}
	}
}

func IsAdmin() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenService := service.TokenService{}
		res, err := tokenService.Verify(c)
		if err != nil {
			log.Println("invalid token")
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"data": err.Error(),
			})
			c.Abort()
		}
		if res.UserType != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"data": err.Error(),
			})
			c.Abort()
		}

	}
}