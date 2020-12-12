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

func CreatePortfolio(c *gin.Context) {
	portfolio := models.Portfolio{}
	err := c.BindJSON(&portfolio)
	if err != nil {
		log.Println("action=CreatePortfolio bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	portfolioRepository := repository.PortfolioRepository{}
	err = portfolioRepository.Create(&portfolio)
	if err != nil {
		log.Println("action=CreatePortfolio failed to create portfolio")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func UpdatePortfolio(c *gin.Context) {
	portfolio := models.Portfolio{}
	err := c.BindJSON(&portfolio)
	if err != nil {
		log.Println("action=UpdatePortfolio bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	portfolioRepository := repository.PortfolioRepository{}
	err = portfolioRepository.Update(&portfolio)
	if err != nil {
		log.Println("action=UpdatePortfolio failed to update portfolio")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func DeletePortfolio(c *gin.Context) {
	paramPortfolioId , err := strconv.ParseInt(c.Param("portfolioId"), 10, 64)
	if err != nil {
		log.Println("action=DeletePortfolio user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	portfolioRepository := repository.PortfolioRepository{}
	err = portfolioRepository.Delete(paramPortfolioId)
	if err != nil {
		log.Println("action=DeletePortfolio failed to delete portfolio")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func GetUserPortfolios(c *gin.Context) {
	tokenService := service.TokenService{}
	user, err := tokenService.Verify(c)
	if err != nil {
		log.Println("action=GetUserPortfolios user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	portfolioRepository := repository.PortfolioRepository{}
	portfolios, err := portfolioRepository.GetByUserId(user.UserId)
	if err != nil {
		log.Println("action=GetUserPortfolios failed to get portfolio")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *portfolios,
	})

}

func GetUserPortfolioById(c * gin.Context) {
	portfolioId, err := strconv.ParseInt(c.Param("portfolioId"), 10, 64)
	if err != nil {
		log.Println("action=GetUserPortfolioById user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	portfolioRepository := repository.PortfolioRepository{}
	portfolios, err := portfolioRepository.GetByPortfolioId(portfolioId)
	if err != nil {
		log.Println("action=GetUserPortfolioById failed to get portfolios")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *portfolios,
	})
}

func GetAllPortfolios(c *gin.Context) {
	portfolios := &[]models.Portfolio{}
	portfolioRepository := repository.PortfolioRepository{}
	portfolios, err := portfolioRepository.GetAll()
	if err != nil {
		log.Println("action=GetAllPortfolios bizpack is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *portfolios,
	})

}

func PortfolioRouter(group *gin.RouterGroup){
	myPageEngine := group.Group("/mypage")
	myPageEngine.Use(middleware.IsLogin())
	{
		BizpackEngine := myPageEngine.Group("/portfolio")
		{
			BizpackEngine.POST("/create", CreatePortfolio)
			BizpackEngine.PUT("/:portfolioId/update", UpdatePortfolio)
			BizpackEngine.DELETE("/:portfolioId/delete", DeletePortfolio)
			BizpackEngine.GET("/get/:portfolioId", GetUserPortfolioById)
			BizpackEngine.GET("/all", GetAllPortfolios)
			BizpackEngine.GET("/", GetUserPortfolios)
		}
	}
}

