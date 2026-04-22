package services

import (
	"context"
	"fmt"
	"github.com/swastiijain24/psp/internals/httpclient"
	repo "github.com/swastiijain24/psp/internals/repositories"
	"github.com/swastiijain24/psp/internals/utils"
)

type AccountService interface {
	DiscoverAccounts(ctx context.Context, phone string, bankCode string) ([]string, error)
	LinkAccount(ctx context.Context, vpaId, accountId string, bankCode string) error
	GetTransactionHistory(ctx context.Context, vpaId string) ([]repo.Transaction, error)
	SetMpin(ctx context.Context, vpaId string, Mpin string) error
	ChangeMpin(ctx context.Context, vpaId string, oldMpinEn string ) error
	GetBalance(ctx context.Context, vpaId string, MpinEn string) (string, error)
}

type Accountsvc struct {
	npciClient         httpclient.Client
	transactionService TransactionService
}

func NewAccountService(npciClient httpclient.Client, transactionService TransactionService) AccountService {
	return &Accountsvc{
		npciClient:         npciClient,
		transactionService: transactionService,
	}
}

func (s *Accountsvc) DiscoverAccounts(ctx context.Context, phone string, bankCode string) ([]string, error) {

	cleanPhone, err := utils.ValidatePhoneNumber(phone)
	if err != nil {
		return []string{}, err
	}

	accountIds, err := s.npciClient.DiscoverAccounts(ctx, cleanPhone, bankCode)
	if err != nil {
		return []string{}, err
	}

	return accountIds, nil
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

func (s *Accountsvc) SetMpin(ctx context.Context, vpaId string, mpinEn string) error {
	err := s.npciClient.SetMpin(ctx, vpaId, mpinEn)
	if err != nil {
		return err
	}
	return nil
}

func (s *Accountsvc) ChangeMpin(ctx context.Context, vpaId string, oldMpinEn string) error {
	err := s.npciClient.ChangeMpin(ctx, vpaId, oldMpinEn)
	if err != nil {
		return err
	}
	return nil
}

func (s *Accountsvc) GetBalance(ctx context.Context, vpaId string, mpinEn string) (string, error) {
	paise, err := s.npciClient.GetBalance(ctx, vpaId, mpinEn)
	if err != nil {
		return "", err
	}
	return utils.PaiseToRupees(paise), nil
}

func (s *Accountsvc) GetTransactionHistory(ctx context.Context, vpaId string) ([]repo.Transaction, error) {

	transactions, err := s.transactionService.GetTransactionHistory(ctx, vpaId)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
