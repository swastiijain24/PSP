package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/swastiijain24/psp/internal/services"
)

type AccountHandler struct {
	accountService services.AccountService 
}

func NewAccountHandler(accountService services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (h *AccountHandler) LinkAccount(c *gin.Context) {
	var linkReq LinkAccountReq
	if err := c.ShouldBindJSON(&linkReq); err != nil {
		c.JSON(400, gin.H{"error":err})
	}
	// response := h.accountService.LinkAccount(c.Request.Context(), linkReq.VpaId, linkReq.AccountId, linkReq.BankCode)

	// c.JSON(201, gin.H{"response":response})
}

type LinkAccountReq struct {
	VpaId     string `json:"vpa_id" binding:"required"`
	AccountId string `json:"account_id" binding:"required"`
	BankCode  string `json:"bank_code" binding:"required"`
}