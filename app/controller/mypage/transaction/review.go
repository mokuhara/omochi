package transaction

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"omochi/app/models"
	"omochi/app/repository"
	"omochi/app/service"
	"strconv"
)

func CreateReview(c *gin.Context){
	review := models.Review{}
	err := c.BindJSON(&review)
	if err != nil {
		log.Println("action=CreateReview bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	reviewRepository := repository.ReviewRepository{}
	err = reviewRepository.Create(&review)
	if err != nil {
		log.Println("action=CreateReview failed to create review")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}
func UpdateReview(c *gin.Context){
	review := models.Review{}
	err := c.BindJSON(&review)
	if err != nil {
		log.Println("action=UpdateReview bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	reviewRepository := repository.ReviewRepository{}
	err = reviewRepository.Update(&review)
	if err != nil {
		log.Println("action=UpdateReview failed to update review")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}
func DeleteReview(c *gin.Context){
	paramReviewId, err := strconv.ParseInt(c.Param("reviewId"), 10, 64)
	if err != nil {
		log.Println("action=DeleteReview user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	reviewRepository := repository.ReviewRepository{}
	err = reviewRepository.Delete(paramReviewId)
	if err != nil {
		log.Println("action=DeleteReview failed to delete review")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}
func GetTransactionReviewByUserId(c *gin.Context){
	tokenService := service.TokenService{}
	user, err := tokenService.Verify(c)
	if err != nil {
		log.Println("action=GetTransactionReviewByUserId user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	reviewRepository := repository.ReviewRepository{}
	reviews, err := reviewRepository.GetByUserId(user.UserId)
	if err != nil {
		log.Println("action=GetTransactionReviewByUserId failed to get review")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *reviews,
	})
}

func GetTransactionReviewByTransactionId(c *gin.Context){
	paramTransactionId, err := strconv.ParseInt(c.Param("transactionId"), 10, 64)
	if err != nil {
		log.Println("action=GetTransactionReviewByTransactionId user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	reviewRepository := repository.ReviewRepository{}
	reviews, err := reviewRepository.GetByTransactionId(paramTransactionId)
	if err != nil {
		log.Println("action=GetTransactionReviewByTransactionId failed to get review")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *reviews,
	})
}
