package mypage

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"omochi/app/models"
	"omochi/app/repository"
	"omochi/app/service"
	"omochi/middleware"
	"strconv"
)

func CreateBizpack(c *gin.Context) {
	bizpack := models.Bizpack{}
	err := c.BindJSON(&bizpack)
	if err != nil {
		log.Println("action=CreateBizpack bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	bizpackRepository := repository.BizpackRepository{}
	err = bizpackRepository.Create(&bizpack)
	if err != nil {
		log.Println("action=CreateBizpack failed to create bizpack")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func UpdateBizpack(c *gin.Context) {
	bizpack := models.Bizpack{}
	err := c.BindJSON(&bizpack)
	if err != nil {
		log.Println("action=UpdateBizpack bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	bizpackRepository := repository.BizpackRepository{}
	err = bizpackRepository.Update(&bizpack)
	if err != nil {
		log.Println("action=UpdateBizpack failed to update bizpack")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func DeleteBizpack(c *gin.Context) {
	paramBizpackId , err := strconv.ParseInt(c.Param("bizpackId"), 10, 64)
	if err != nil {
		log.Println("action=DeleteBizpack user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	bizpackRepository := repository.BizpackRepository{}
	err = bizpackRepository.Delete(paramBizpackId)
	if err != nil {
		log.Println("action=DeleteBizpack failed to delete bizpack")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func GetUserBizpacks(c *gin.Context) {
	tokenService := service.TokenService{}
	user, err := tokenService.Verify(c)
	if err != nil {
		log.Println("action=DeleteBizpack user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	bizpackRepository := repository.BizpackRepository{}
	bizpacks, err := bizpackRepository.GetByUserId(user.UserId)
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

func GetAllBizpacks(c *gin.Context) {
	bizpacks := &[]models.Bizpack{}
	bizpackRepository := repository.BizpackRepository{}
	bizpacks, err := bizpackRepository.GetAll()
	if err != nil {
		log.Println("action=GetAllBizpacks bizpack is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *bizpacks,
	})

}

func BizPackRouter(group *gin.RouterGroup){
	myPageEngine := group.Group("/mypage")
	myPageEngine.Use(middleware.IsLogin())
	{
		BizpackEngine := myPageEngine.Group("/bizpack")
		{
			BizpackEngine.GET("/", GetUserBizpacks)
			BizpackEngine.GET("/all", GetAllBizpacks)
			BizpackEngine.POST("/create", CreateBizpack)
			BizpackEngine.PUT("/:bizpackId/update", UpdateBizpack)
			BizpackEngine.DELETE("/:bizpackId/delete", DeleteBizpack)
		}
	}
}

