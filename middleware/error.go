package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			log.Print(err.Err)

			data := map[string]string{"error": err.Error()}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": err.Meta,
				"data": data,
			})
		}
	}
}