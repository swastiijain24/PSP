package services

import (
	"context"
	"fmt"

	"github.com/swastiijain24/psp/internal/httpclient"
)

type PaymentService interface{}

type PaymentSvc struct{
	npciClient *httpclient.NpciClient
}

func NewPaymentService(npciClient *httpclient.NpciClient) PaymentService {
	return &PaymentSvc{
		npciClient: npciClient,
	}
}

func (s *PaymentSvc) Pay(ctx context.Context, transactionId string, payerVpa string, payeeVpa string, amount int64, mpin string) error {

	if amount <=0 {
		return fmt.Errorf("invalid amount")
	}

	if payerVpa == payeeVpa{
		return fmt.Errorf("error")
	}

	return nil 
}