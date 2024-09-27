package resolver

import (
	"context"
	"github.com/coretech/dtc-api-gateway/internal/app/gateway/domain"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type TransactionService interface {
	Get(ctx context.Context, filter domain.TransactionFilter, limit int, skip int) ([]domain.Transaction, error)
}

type Resolver struct {
	transaction TransactionService
}

func New(txs TransactionService) *Resolver {
	return &Resolver{
		transaction: txs,
	}
}
