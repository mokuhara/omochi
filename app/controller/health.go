package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "ok",
	})
}
