package mypage

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"omochi/app/models"
	"omochi/app/repository"
	"omochi/middleware"
	"strconv"
)

func GetProfile(c *gin.Context){
	paramUserId, err := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err != nil {
		log.Println("action=GetProfile user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	profileRepository := repository.ProfileRepository{}
	profile, err := profileRepository.Get(paramUserId)
	if err != nil {
		log.Println("action=GetProfile profile is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": profile,
	})
}

func CreateProfile(c *gin.Context){
	profile := models.Profile{}
	err := c.BindJSON(&profile)
	if err != nil {
		log.Println("action=CreateProfile bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	profileRepository := repository.ProfileRepository{}
	err = profileRepository.Create(&profile)
	if err != nil {
		log.Println("action=CreateProfile failed to create profile")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func UpdateProfile(c *gin.Context){
	profile := models.Profile{}
	err := c.BindJSON(&profile)
	if err != nil {
		log.Println("action=UpdateProfile bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	profileRepository := repository.ProfileRepository{}
	oldProfile, err := profileRepository.Get(profile.UserID)
	if err != nil {
		log.Println("action=UpdateProfile profile is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	profile.ID = oldProfile.ID
	err = profileRepository.Update(&profile)
	if err != nil {
		log.Println("action=UpdateProfile failed to update profile")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func DeleteProfile(c *gin.Context) {
	paramUserId, err := strconv.ParseInt(c.Param("userId"), 10,64)
	if err != nil {
		log.Println("action=DeleteProfile user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	profileRepository := repository.ProfileRepository{}
	profile, err := profileRepository.Get(paramUserId)
	if err != nil {
		log.Println("action=DeleteProfile failed to get profile")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	err = profileRepository.Delete(int64(profile.ID))
	if err != nil {
		log.Println("action=DeleteProfile failed to delete profile")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func ProfileRouter(group *gin.RouterGroup){
	myPageEngine := group.Group("/mypage/:userId")
	myPageEngine.Use(middleware.IsLogin())
	{
		UserProfileEngine := myPageEngine.Group("/user")
		{
			UserProfileEngine.GET("/", GetProfile)
			UserProfileEngine.POST("/create", CreateProfile)
			UserProfileEngine.PUT("/update", UpdateProfile)
			UserProfileEngine.DELETE("/delete", DeleteProfile)
		}
	}
}