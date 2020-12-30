package transaction

import "github.com/gin-gonic/gin"

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
			reviewEngine.GET("/:transactionId", GetTransactionReviewByTransactionId)
		}
		videoMeetingEngine := myPageEngine.Group("/videomeeting")
		{
			videoMeetingEngine.POST("/create", CreateVideoMeeting)
			videoMeetingEngine.PUT("/:videoMeetingId/update", UpdateVideoMeeting)
			videoMeetingEngine.DELETE("/:videoMeetingId/delete", DeleteVideoMeeting)
			videoMeetingEngine.GET("/:transactionId", GetVideoMeetingByTransactionId)
		}
	}
}