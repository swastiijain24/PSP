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
	err := h.accountService.LinkAccount(c.Request.Context(), linkReq.VpaId, linkReq.AccountId, linkReq.BankCode)
	if err != nil {
		c.JSON(400, gin.H{"error":err})
	}
	c.JSON(201, gin.H{"response":"Account linked successfully"})
}

func (h *AccountHandler) CreateAccount(c *gin.Context){
	var accReq CreateAccountReq 
	if err := c.ShouldBindJSON(&accReq); err != nil{
		c.JSON(400, gin.H{"error":err})
	}
	account, err := h.accountService.CreateAccount(c.Request.Context(), accReq.Name, accReq.Phone, accReq.Mpin)
	if err != nil {
		c.JSON(400, gin.H{"error":err})
	} 
	
	c.JSON(201, gin.H{"account":account})

}

type LinkAccountReq struct {
	VpaId     string `json:"vpa_id" binding:"required"`
	AccountId string `json:"account_id" binding:"required"`
	BankCode  string `json:"bank_code" binding:"required"`
}


type CreateAccountReq struct {
    Name  string `json:"name" binding:"required,min=1,max=255"`
    Phone string `json:"phone" binding:"required,e164"`  
	Mpin string `json:"mpin_hash" binding:"required"`
}