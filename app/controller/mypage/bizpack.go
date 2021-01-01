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

var bizpackRepository = repository.BizpackRepository{}

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

func Show(c *gin.Context) {
	user, err := authenticate(c)

	if err != nil {
		log.Println("action=Show user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)

		return
	}

	bizpackId, err := strconv.ParseInt(c.Param("bizpackId"), 10, 64)

	if err != nil {
		log.Println("action=ShowBizpack bizpackId is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)

		return
	}

	bizpack, err := bizpackRepository.GetByUserIDAndBizpackId(user.UserId, bizpackId)

	if err != nil {
		log.Println("action=GetByUserIDAndBizpackId failed to get bizpack")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *bizpack,
	})
}

func Update(c *gin.Context) {
	user, err := authenticate(c)

	if err != nil {
		log.Println("action=Update user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)

		return
	}

	bizpackId, err := strconv.ParseInt(c.Param("bizpackId"), 10, 64)

	if err != nil {
		log.Println("action=Update bizpackId is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)

		return
	}

	bizpack, err := bizpackRepository.GetByUserIDAndBizpackId(user.UserId, bizpackId)

	if err != nil {
		log.Println("action=GetByUserIDAndBizpackId failed to get bizpack")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)

		return
	}

	bindErr := c.BindJSON(bizpack)

	if bindErr != nil {
		log.Println("action=UpdateBizpack bind error")
		c.Error(bindErr).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)

		return
	}

	err = bizpackRepository.Update(bizpack)

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
	user, err := authenticate(c)

	if err != nil {
		log.Println("action=DeleteBizpack user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)

		return
	}

	// bizpackIdをURLから取得する
	bizpackId, err := strconv.ParseInt(c.Param("bizpackId"), 10, 64)

	if err != nil {
		log.Println("action=DeleteBizpack user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)

		return
	}

	// UserIDとbizpackIdで削除対象のレコードを検索する
	var deletable = bizpackRepository.CheckUserBizpack(user.UserId, bizpackId)

	// 削除を実行する
	if deletable {
		err = bizpackRepository.Delete(bizpackId)

		if err != nil {
			log.Println("action=DeleteBizpack failed to delete bizpack")
			c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": deletable,
	})
}

func BizPackRouter(group *gin.RouterGroup){
	myPageEngine := group.Group("/mypage")
	myPageEngine.Use(middleware.IsLogin())
	{
		BizpackEngine := myPageEngine.Group("/bizpacks")
		{
			BizpackEngine.GET("", Index)
			BizpackEngine.POST("", Create)
			BizpackEngine.GET("/:bizpackId", Show)
			BizpackEngine.PUT("/:bizpackId", Update)
			BizpackEngine.DELETE("/:bizpackId", Delete)
		}
	}
}
