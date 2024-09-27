package hdm_go_transactions_api

import (
	"context"
	"github.com/coretech/dtc-api-gateway/internal/app/gateway/domain"
	"github.com/coretech/dtc-api-gateway/pkg/client/hdm-go-transactions-api/genqlient"
	"github.com/coretech/dtc-api-gateway/pkg/client/hdm-go-transactions-api/gqlgenc"
)

var (
	_ Client = (*genqlient.Client)(nil)
	_ Client = (*gqlgenc.Client)(nil)
)

type Client interface {
	GetTransactions(ctx context.Context, filter domain.TransactionFilter, limit int, skip int) ([]domain.Transaction, error)
}
