package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func IsLogin() gin.HandlerFunc{
	return func(c *gin.Context){
		paramUserId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	}
}