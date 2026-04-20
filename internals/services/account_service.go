package services

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/swastiijain24/psp/internals/httpclient"
	"github.com/swastiijain24/psp/internals/utils"
)

type AccountService interface {
	LinkAccount(ctx context.Context, vpaId, accountId string, bankCode string) error
	CreateAccount(ctx context.Context, name string, phone string, mpin string) error
}

type Accountsvc struct {
	npciClient *httpclient.NpciClient
}

func NewAccountService(npciClient *httpclient.NpciClient) AccountService {
	return &Accountsvc{
		npciClient: npciClient,
	}
}

func (s *Accountsvc) LinkAccount(ctx context.Context, vpaId string, accountId string, bankCode string) error {

	err := utils.ValidateVPA(vpaId)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	if err = s.npciClient.LinkAccount(ctx, vpaId, accountId, bankCode); err != nil {
		return err
	}
	return nil
}

func (s *Accountsvc) CreateAccount(ctx context.Context, name string, phone string, mpin string) error {
	err := utils.ValidateMPIN(mpin)
	if err != nil {
		return err
	}

	if name == "" {
		return fmt.Errorf("name cannot be null")
	}

	cleanPhone, err := utils.ValidatePhoneNumber(phone)
	if err != nil {
		return err
	}

	encryptedPin, err := utils.EncryptAES(mpin, []byte(os.Getenv("MPIN_ENCRYPTION_KEY")))
	if err != nil {
		return err
	}

	err = s.npciClient.CreateAccount(ctx, name, cleanPhone, encryptedPin)
	if err != nil {
		return err
	}

	return nil
}

func (s *Accountsvc) GetBalance() {

}

func (s *Accountsvc) GetTransactionHistory(ctx context.Context, accountId string) {
	id := c.Param("id")

	transactions, err := h.

}
