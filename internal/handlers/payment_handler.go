package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/swastiijain24/psp/internal/services"
)

type PaymentHandler struct {
	paymentService *services.PaymentService
}

func NewPaymentHandler(paymentService *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{}
}

func (h *PaymentHandler) Pay(c *gin.Context){
	var params paymentParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.paymentService.Pay(c.Request.Context(), params.TransactionID, params.PayerVPA, params.PayeeVPA, params.Amount, params.Mpin)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}

type paymentParams struct {
	TransactionID string `json:"transaction_id" binding:"required"`
	PayerVPA      string `json:"payer_vpa" binding:"required"`
	PayeeVPA      string `json:"payee_vpa" binding:"required"`
	Amount        int64  `json:"amount" binding:"required"`
	Mpin          string `json:"mpin" binding:"required"`
}