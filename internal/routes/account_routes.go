package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swastiijain24/psp/internal/handlers"
)

func RegisterAccountRoutes(r *gin.Engine, accountHandler *handlers.AccountHandler) {
	accountRoutes := r.Group("/account")
	{
		accountRoutes.POST("/create-account", accountHandler.CreateAccount)
		accountRoutes.POST("/link-account", accountHandler.LinkAccount)
		accountRoutes.GET("balance/:id", accountHandler.GetBalance)
		accountRoutes.GET("transactions/:id", accountHandler.GetTransactions) 
	}
} 

