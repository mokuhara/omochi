package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"omochi/app/service"
	"strconv"
)

func IsLogin() gin.HandlerFunc{
	return func(c *gin.Context){
		paramUserId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		tokenService := service.TokenService{}
		res, err := tokenService.Verify(c)
		if err != nil {
			log.Println("invalid token")
			c.Error(err).SetType(gin.ErrorTypePublic)
			c.Abort()
		}
		if res.UserId != paramUserId {
			c.Error(fmt.Errorf("invalid userId")).SetType(gin.ErrorTypePublic)
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
			c.Error(err).SetType(gin.ErrorTypePublic)
			c.Abort()
		}
		if res.UserType != 0 {
			c.Error(fmt.Errorf("insufficient authority")).SetType(gin.ErrorTypePublic)
			c.Abort()
		}

	}
}