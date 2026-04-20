package routes

import "github.com/gin-gonic/gin"

func RegisterPaymentRoutes(r *gin.Engine) {
	paymentRoutes := r.Group("/payment")
	{
		paymentRoutes.POST("/oay")
	}
}