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

func CreatePayment(c *gin.Context){
	payment := models.Payment{}
	err := c.BindJSON(&payment)
	if err != nil {
		log.Println("action=CreatePayment bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	paymentRepository := repository.PaymentRepository{}
	err = paymentRepository.Create(&payment)
	if err != nil {
		log.Println("action=CreatePayment failed to create payment")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}
func UpdatePayment(c *gin.Context){
	payment := models.Payment{}
	err := c.BindJSON(&payment)
	if err != nil {
		log.Println("action=UpdatePayment bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	paymentRepository := repository.PaymentRepository{}
	err = paymentRepository.Update(&payment)
	if err != nil {
		log.Println("action=UpdatePayment failed to update payment")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}
func DeletePayment(c *gin.Context){
	paramPaymentId, err := strconv.ParseInt(c.Param("paymentId"), 10, 64)
	if err != nil {
		log.Println("action=DeletePayment user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	paymentRepository := repository.PaymentRepository{}
	err = paymentRepository.Delete(paramPaymentId)
	if err != nil {
		log.Println("action=DeletePayment failed to delete payment")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}
func GetTransactionPayment(c *gin.Context){
	tokenService := service.TokenService{}
	user, err := tokenService.Verify(c)
	if err != nil {
		log.Println("action=GetTransactionPayment user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	paymentRepository := repository.PaymentRepository{}
	payments, err := paymentRepository.GetByUserId(user.UserId)
	if err != nil {
		log.Println("action=GetTransactionPayment failed to get payment")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *payments,
	})
}
