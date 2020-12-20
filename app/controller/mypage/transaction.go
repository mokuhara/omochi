package mypage

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"omochi/app/models"
	"omochi/app/repository"
	"omochi/app/service"
	"omochi/config"
	"strconv"
)

func CreateTransaction(c *gin.Context) {
	transaction := models.Transaction{}
	err := c.BindJSON(&transaction)
	if err != nil {
		log.Println("action=CreateTransaction bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
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
func GetTransactionReview(c *gin.Context){
	tokenService := service.TokenService{}
	user, err := tokenService.Verify(c)
	if err != nil {
		log.Println("action=GetTransactionReview user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	reviewRepository := repository.ReviewRepository{}
	reviews, err := reviewRepository.GetByUserId(user.UserId)
	if err != nil {
		log.Println("action=GetTransactionReview failed to get review")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *reviews,
	})
}

func CreateVideoMeeting(c *gin.Context){
	type setting struct {
		Use_pmi string `json:"use_pmi"`
	}

	type zoomCreateRoom struct {
		Topic string  `json:"topic"`
		Type string `json:"type"`
		Start_time string `json:"start_time"`
		Timezone string `json:"timezone"`
		Settings setting
	}

	zoomCreateRoomStruct := zoomCreateRoom{}
	err := c.BindJSON(&zoomCreateRoomStruct)
	if err != nil {
		log.Println("action=CreateVideoMeeting bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
    endpoint := config.Config.ZoomEndpoint
    jwt := config.Config.ZoomJwt
	zoomCreateRoomJson, err := json.Marshal(zoomCreateRoomStruct)
	if err != nil {
		log.Println("action=CreateVideoMeeting zoomCreateRoom Marshal error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	req, _ := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(zoomCreateRoomJson))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + jwt)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("action=CreateVideoMeeting failed to create zoom room")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"data": "failed to create zoom room",
		})
	}

	body, _ := ioutil.ReadAll(resp.Body)
	videoMeeting := models.VideoMeeting{}
	err = json.Unmarshal(body, &videoMeeting)
	if err != nil {
		log.Println("action=CreateVideoMeeting failed to unmarshal videMeeting")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	videoMeetingRepository := repository.VideoMeetingRepository{}
	err = videoMeetingRepository.Create(&videoMeeting)
	if err != nil {
		log.Println("action=CreateVideoMeeting failed to create videoMeeting")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": videoMeeting,
	})
}

func UpdateVideoMeeting(c *gin.Context){
	videoMeeting := models.VideoMeeting{}
	err := c.BindJSON(&videoMeeting)
	if err != nil {
		log.Println("action=UpdateVideoMeeting bind error")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	videoMeetingRepository := repository.VideoMeetingRepository{}
	err = videoMeetingRepository.Update(&videoMeeting)
	if err != nil {
		log.Println("action=UpdateVideoMeeting failed to update videoMeeting")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}
func DeleteVideoMeeting(c *gin.Context){
	paramVideoMeetingId, err := strconv.ParseInt(c.Param("videoMeetingId"), 10, 64)
	if err != nil {
		log.Println("action=DeleteVideoMeeting user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	videoMeetingRepository := repository.VideoMeetingRepository{}
	err = videoMeetingRepository.Delete(paramVideoMeetingId)
	if err != nil {
		log.Println("action=DeleteVideoMeeting failed to delete videoMeeting")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": "",
	})
}
func GetTransactionVideoMeeting(c *gin.Context){
	tokenService := service.TokenService{}
	user, err := tokenService.Verify(c)
	if err != nil {
		log.Println("action=GetTransactionVideoMeeting user_id is not found")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusBadRequest)
		return
	}
	videoMeetingRepository := repository.VideoMeetingRepository{}
	videoMeetings, err := videoMeetingRepository.GetByUserId(user.UserId)
	if err != nil {
		log.Println("action=GetTransactionVideoMeeting failed to get videoMeetings")
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": *videoMeetings,
	})
}

func TransactionRouter(group *gin.RouterGroup) {
	myPageEngine := group.Group("/mypage")
	//myPageEngine.Use(middleware.IsLogin())
	{
		transactionEngine := myPageEngine.Group("/transaction")
		{
			transactionEngine.POST("/create", CreateTransaction)
			transactionEngine.PUT("/:transactionId/update", UpdateTransaction)
			transactionEngine.DELETE("/:transactionId/delete", DeleteTransaction)
			transactionEngine.GET("/", GetUserTransaction)
		}
		paymentEngine := myPageEngine.Group("/payment")
		{
			paymentEngine.POST("/create", CreatePayment)
			paymentEngine.PUT("/:paymentId/update", UpdatePayment)
			paymentEngine.DELETE("/:paymentId/delete", DeletePayment)
			paymentEngine.GET("/", GetTransactionPayment)
		}
		reviewEngine := myPageEngine.Group("/review")
		{
			reviewEngine.POST("/create", CreateReview)
			reviewEngine.PUT("/:reviewId/update", UpdateReview)
			reviewEngine.DELETE("/:reviewId/delete", DeleteReview)
			reviewEngine.GET("/", GetTransactionReview)
		}
		videoMeetingEngine := myPageEngine.Group("/videomeeting")
		{
			videoMeetingEngine.POST("/create", CreateVideoMeeting)
			videoMeetingEngine.PUT("/:videoMeetingId/update", UpdateVideoMeeting)
			videoMeetingEngine.DELETE("/:videoMeetingId/delete", DeleteVideoMeeting)
			videoMeetingEngine.GET("/", GetTransactionVideoMeeting)
		}
	}
}