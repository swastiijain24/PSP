package services

import (
	"context"

	repo "github.com/swastiijain24/psp/internals/repositories"
)

type TransactionService interface{

}

type TransactionSvc struct{
	repo repo.Querier
}

func NewTransactionService(repo repo.Querier) TransactionService {
	return &TransactionSvc{
		repo: repo,
	}
}

func GetTransactionHistory(ctx context.Context, accountId string) ([]repo.Transaction, error){
	
}