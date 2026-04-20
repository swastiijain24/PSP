package services

import (
	"context"
	"fmt"

	"github.com/swastiijain24/psp/internal/httpclient"
	"github.com/swastiijain24/psp/internal/utils"
)

type AccountService interface {
	LinkAccount(ctx context.Context, vpaId, accountId string, bankCode string) error 
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

	if err = s.npciClient.LinkAccount(ctx, vpaId, accountId, bankCode); err !=nil {
		return err  
	}
	return nil 
}
