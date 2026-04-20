package routes

import "github.com/gin-gonic/gin"

func RegisterAccountRoutes(r *gin.Engine) {
	accountRoutes := r.Group("/account")
	{
		accountRoutes.POST("/link-account", accountHandler.LinkAccount)
		accountRoutes.GET("balance/:id", accountHandler.GetBalance)
		accountRoutes.GET("transactions/:id", accountHandler.GetTransactions) 
	}
} 

