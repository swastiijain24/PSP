package services

import (
	"context"

	repo "github.com/swastiijain24/psp/internals/repositories"
)

type TransactionService interface{
	GetTransactionHistory(ctx context.Context, vpdaId string) ([]repo.Transaction, error)
}

type TransactionSvc struct{
	repo repo.Querier
}

func NewTransactionService(repo repo.Querier) TransactionService {
	return &TransactionSvc{
		repo: repo,
	}
}

func (s* TransactionSvc) GetTransactionHistory(ctx context.Context, vpdaId string) ([]repo.Transaction, error){
	return s.repo.GetTransactionHistory(ctx, repo.GetTransactionHistoryParams{
		PayerVpa: vpdaId,
		Limit: 10,
		Offset: 20,
	})
}