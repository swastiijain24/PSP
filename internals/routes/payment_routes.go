package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swastiijain24/psp/internals/handlers"
)

func RegisterPaymentRoutes(r *gin.Engine, paymentHandler *handlers.PaymentHandler) {
	paymentRoutes := r.Group("/payment")
	{
		paymentRoutes.POST("/pay", paymentHandler.Pay)
		paymentRoutes.GET("/status/:txnid", paymentHandler.GetStatus)	
	}
}
