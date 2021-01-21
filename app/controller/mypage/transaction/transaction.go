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

func authenticate(c *gin.Context) (auth *service.Auth, err error) {
	tokenService := service.TokenService{}
	user, err := tokenService.Verify(c)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// 一旦client側からの作成のみを許容
func CreateTransaction(c *gin.Context) {
	user, authErr := authenticate(c)

	if authErr != nil {
		log.Println("action=DeleteBizpack user_id is not found")
		c.Error(authErr).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)

		return
	}

	transaction := models.Transaction{}
	err := c.BindJSON(&transaction)

	if err != nil {
		log.Println("action=CreateTransaction bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}

	transaction.ClientUserID = user.UserId

	transactionRepository := repository.TransactionRepository{}
	err = transactionRepository.Create(&transaction)

	if err != nil {
		log.Println("action=CreateTransaction failed to create transaction")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func UpdateTransaction(c *gin.Context) {
	transaction := models.Transaction{}
	err := c.BindJSON(&transaction)
	if err != nil {
		log.Println("action=UpdateTransaction bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	transactionRepository := repository.TransactionRepository{}
	err = transactionRepository.Update(&transaction)
	if err != nil {
		log.Println("action=UpdateTransaction failed to update transaction")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func DeleteTransaction(c *gin.Context) {
	paramTransactionId, err := strconv.ParseInt(c.Param("transactionId"), 10, 64)
	if err != nil {
		log.Println("action=DeleteTransaction user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	transactionRepository := repository.TransactionRepository{}
	err = transactionRepository.Delete(paramTransactionId)
	if err != nil {
		log.Println("action=DeleteTransaction failed to delete bizpack")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}

func GetUserTransaction(c *gin.Context) {
	tokenService := service.TokenService{}
	user, err := tokenService.Verify(c)
	if err != nil {
		log.Println("action=GetUserTransaction user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	transactionRepository := repository.TransactionRepository{}
	transactions, err := transactionRepository.GetByUserId(user.UserId)
	if err != nil {
		log.Println("action=GetUserTransaction failed to get transaction")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *transactions,
	})
}