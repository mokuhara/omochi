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

func Show(c *gin.Context) {
	bizpackId, err := strconv.ParseInt(c.Param("bizpackId"), 10, 64)

	if err != nil {
		log.Println("action=DeleteBizpack user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}

	bizpackRepository := repository.BizpackRepository{}
	bizpacks, err := bizpackRepository.GetByBizpackId(bizpackId)

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

func Create(c *gin.Context) {
	user, err := authenticate(c)

	if err != nil {
		log.Println("action=DeleteBizpack user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}

	bizpack := models.Bizpack{}
	bindErr := c.BindJSON(&bizpack)

	if bindErr != nil {
		log.Println("action=CreateBizpack bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}

	// 他のユーザと紐づかせないためbindの後にセットする
	bizpack.UserID = user.UserId

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

func Update(c *gin.Context) {
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

func Delete(c *gin.Context) {
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

func BizPackRouter(group *gin.RouterGroup){
	myPageEngine := group.Group("/mypage")
	myPageEngine.Use(middleware.IsLogin())
	{
		BizpackEngine := myPageEngine.Group("/bizpacks")
		{
			BizpackEngine.POST("", Create)
			BizpackEngine.PUT("/:bizpackId", Update)
			BizpackEngine.DELETE("/:bizpackId", Delete)
			BizpackEngine.GET("/:bizpackId", Show)
			BizpackEngine.GET("", Index)
		}
	}
}
