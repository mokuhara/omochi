package specialist

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"omochi/app/repository"
)

func GetAllProfile(c *gin.Context){
	profileRepository := repository.ProfileRepository{}
	profiles, err := profileRepository.GetAll()
	if err != nil {
		log.Println("action=GetAllProfile failed to get profiles")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": profiles,
	})
}

func Router(group *gin.RouterGroup){
	specialistEngine := group.Group("/specialist")
	{
		specialistEngine.GET("/", GetAllProfile)
	}
}
