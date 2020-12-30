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

func CreateIssue(c *gin.Context) {
	issue := models.Issue{}
	err := c.BindJSON(&issue)
	if err != nil {
		log.Println("action=CreateIssue bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	issueRepository := repository.IssueRepository{}
	err = issueRepository.Create(&issue)
	if err != nil {
		log.Println("action=CreateIssue failed to create issue")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func UpdateIssue(c *gin.Context) {
	issue := models.Issue{}
	err := c.BindJSON(&issue)
	if err != nil {
		log.Println("action=UpdateIssue bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	issueRepository := repository.IssueRepository{}
	err = issueRepository.Update(&issue)
	if err != nil {
		log.Println("action=UpdateIssue failed to update issue")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func DeleteIssue(c *gin.Context) {
	paramIssueId , err := strconv.ParseInt(c.Param("issueId"), 10, 64)
	if err != nil {
		log.Println("action=DeleteIssue user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	issueRepository := repository.IssueRepository{}
	err = issueRepository.Delete(paramIssueId)
	if err != nil {
		log.Println("action=DeleteIssue failed to delete issue")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func GetUserIssues(c *gin.Context) {
	tokenService := service.TokenService{}
	user, err := tokenService.Verify(c)
	if err != nil {
		log.Println("action=GetUserIssues user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	issueRepository := repository.IssueRepository{}
	issues, err := issueRepository.GetByUserId(user.UserId)
	if err != nil {
		log.Println("action=GetUserIssues failed to get issue")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *issues,
	})

}

func GetUserIssueById(c * gin.Context) {
	issueId, err := strconv.ParseInt(c.Param("issueId"), 10, 64)
	if err != nil {
		log.Println("action=GetUserIssueById user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	issueRepository := repository.IssueRepository{}
	issues, err := issueRepository.Find(issueId)
	if err != nil {
		log.Println("action=GetUserIssueById failed to get issue")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *issues,
	})
}

func GetAllIssues(c *gin.Context) {
	issues := &[]models.Issue{}
	issueRepository := repository.IssueRepository{}
	issues, err := issueRepository.GetAll()
	if err != nil {
		log.Println("action=GetAllIssues issue is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *issues,
	})

}

func IssueRouter(group *gin.RouterGroup){
	myPageEngine := group.Group("/mypage")
	myPageEngine.Use(middleware.IsLogin())
	{
		BizpackEngine := myPageEngine.Group("/issue")
		{
			BizpackEngine.POST("/create", CreateIssue)
			BizpackEngine.PUT("/:issueId/update", UpdateIssue)
			BizpackEngine.DELETE("/:issueId/delete", DeleteIssue)
			BizpackEngine.GET("/get/:issueId", GetUserIssueById)
			BizpackEngine.GET("/all", GetAllIssues)
			BizpackEngine.GET("/", GetUserIssues)
		}
	}
}