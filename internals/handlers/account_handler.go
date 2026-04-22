package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/swastiijain24/psp/internals/services"
)

type AccountHandler struct {
	accountService services.AccountService
}

func NewAccountHandler(accountService services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (h *AccountHandler) Discover(c *gin.Context) {
	var accReq AccountReq
	if err := c.ShouldBindJSON(&accReq); err != nil {
		c.JSON(400, gin.H{"error": err})
	}
	accounts, err := h.accountService.DiscoverAccounts(c.Request.Context(), accReq.Phone, accReq.BankCode)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
	}

	c.JSON(201, gin.H{"accounts": accounts})
}

func (h *AccountHandler) Link(c *gin.Context) {
	var linkReq LinkAccountReq
	if err := c.ShouldBindJSON(&linkReq); err != nil {
		c.JSON(400, gin.H{"error": err})
	}
	err := h.accountService.LinkAccount(c.Request.Context(), linkReq.VpaId, linkReq.AccountId, linkReq.BankCode)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
	}
	c.JSON(201, gin.H{"response": "Account linked successfully"})
}

func (h *AccountHandler) SetMpin(c *gin.Context) {
	vpaId := c.Param("vpaId")

	var Mpin string
	if err := c.ShouldBindJSON(&Mpin); err != nil {
		c.JSON(400, gin.H{"error": err})
	}

	err := h.accountService.SetMpin(c.Request.Context(), vpaId, Mpin)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
	}
	c.JSON(200, gin.H{"response": "Mpin set successfully"})
}

func (h *AccountHandler) ChangeMpin(c *gin.Context) {
	vpaId := c.Param("vpaId")

	var mpins Mpins
	if err := c.ShouldBindJSON(&mpins); err != nil {
		c.JSON(400, gin.H{"error": err})
	}

	err := h.accountService.ChangeMpin(c.Request.Context(), vpaId, mpins.OldMpin, mpins.NewMpin)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
	}
	c.JSON(200, gin.H{"response": "Mpin set successfully"})
}

func (h *AccountHandler) GetTransactionHistory(c *gin.Context) {
	vpaId := c.Param("vpaId")
	transactions, err := h.accountService.GetTransactionHistory(c.Request.Context(), vpaId)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
	}
	c.JSON(200, transactions)
}

func (h *AccountHandler) GetBalance(c *gin.Context) {
	vpaId := c.Param("vpaId")

	var Mpin string
	if err := c.ShouldBindJSON(&Mpin); err != nil {
		c.JSON(400, gin.H{"error": err})
	}
	balance, err := h.accountService.GetBalance(c.Request.Context(), vpaId, Mpin)
	if err != nil {
		c.JSON(400, gin.H{"error": "error fetching account balance"})
	}
	c.JSON(200, gin.H{"balance": balance})
}

type LinkAccountReq struct {
	VpaId     string `json:"vpa_id" binding:"required"`
	AccountId string `json:"account_id" binding:"required"`
	BankCode  string `json:"bank_code" binding:"required"`
}

type AccountReq struct {
	Phone    string `json:"phone" binding:"required,e164"`
	BankCode string `json:"bank_code" binding:"required"`
}

type Mpins struct{
	OldMpin string `json:"old_mpin" binding:"required"`
	NewMpin string `json:"new_mpin" binding:"required"`
}