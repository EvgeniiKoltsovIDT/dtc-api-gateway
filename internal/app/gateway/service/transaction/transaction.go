package transaction

import (
	"context"
	"github.com/coretech/dtc-api-gateway/internal/app/gateway/domain"
)

type TransactionsClient interface {
	GetTransactions(ctx context.Context, filter domain.TransactionFilter, limit int, skip int) ([]domain.Transaction, error)
}

type Service struct {
	transactionsClient TransactionsClient
}

func New(transactionsClient TransactionsClient) *Service {
	return &Service{
		transactionsClient: transactionsClient,
	}
}
