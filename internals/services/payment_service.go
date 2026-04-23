package services

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/swastiijain24/psp/internals/httpclient"
	repo "github.com/swastiijain24/psp/internals/repositories"
	"github.com/swastiijain24/psp/internals/utils"
)

type PaymentService interface {
	Pay(ctx context.Context, payerVpa string, payeeVpa string, amount string, mpin string, remarks string) error
	GetStatus(ctx context.Context, transactionId string) (string , error)
}

type PaymentSvc struct {
	npciClient httpclient.Client
	repo       repo.Querier
}

func NewPaymentService(repo repo.Querier, npciClient httpclient.Client) PaymentService {
	return &PaymentSvc{
		repo:       repo,
		npciClient: npciClient,
	}
}

func (s *PaymentSvc) Pay(ctx context.Context, payerVpa string, payeeVpa string, amount string, mpin string, remarks string) error {

	num, err := strconv.Atoi(amount)
	if err != nil {
		return err
	}

	if num <= 0 {
		return fmt.Errorf("invalid amount")
	}

	if payerVpa == payeeVpa {
		return fmt.Errorf("error")
	}

	transactionId := utils.GenerateSortableTxnID()

	amountInPaise, err := utils.RupeesToPaise(amount)
	if err != nil {
		return err
	}

	mpinEncrypted, err := utils.EncryptAES(mpin, []byte(os.Getenv("MPIN_ENCRYPTION_KEY")))
	if err != nil {
		return err
	}

	_, err = s.repo.CreateTransaction(ctx, repo.CreateTransactionParams{
		TransactionID: transactionId,
		PayerVpa: payerVpa,
		PayeeVpa: payeeVpa,
		Amount: amountInPaise,
		Remarks: utils.ToPGText(remarks),
	})
	if err != nil {
		return err 
	}

	err = s.npciClient.PaymentRequest(ctx, transactionId, payerVpa, payeeVpa, amountInPaise, mpinEncrypted)
	if err != nil {
		return err 
	}
	
	return nil
}

func (s *PaymentSvc) GetStatus(ctx context.Context, transactionId string) (string , error){
	return s.npciClient.GetStatus(ctx, transactionId)
}
