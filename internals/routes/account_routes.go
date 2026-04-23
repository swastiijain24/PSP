package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swastiijain24/psp/internals/handlers"
)

func RegisterAccountRoutes(r *gin.Engine, accountHandler *handlers.AccountHandler) {
	accountRoutes := r.Group("/account")
	{
		accountRoutes.POST("/discover", accountHandler.Discover)
		accountRoutes.POST("/link", accountHandler.Link)
		accountRoutes.POST("/mpin", accountHandler.SetMpin)
		accountRoutes.PUT("/mpin", accountHandler.ChangeMpin)
		accountRoutes.POST("/balance/:vpaId", accountHandler.GetBalance)
		accountRoutes.GET("/transactions/:vpaId?page=1&limit=20", accountHandler.GetTransactionHistory)
	}
}
