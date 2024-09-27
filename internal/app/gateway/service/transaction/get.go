package transaction

import (
	"context"
	"github.com/coretech/dtc-api-gateway/internal/app/gateway/domain"
)

func (s *Service) Get(ctx context.Context, filter domain.TransactionFilter, limit int, skip int) ([]domain.Transaction, error) {
	txs, err := s.transactionsClient.GetTransactions(ctx, filter, limit, skip)
	if err != nil {
		return nil, err
	}

	return txs, nil
}
