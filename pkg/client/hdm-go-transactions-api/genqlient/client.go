package genqlient

import (
	"context"
	"github.com/Khan/genqlient/graphql"
	"github.com/coretech/dtc-api-gateway/internal/app/gateway/domain"
	"net/http"
)

//go:generate go run github.com/Khan/genqlient

type Client struct {
	client graphql.Client
}

func New(baseUrl string) *Client {
	httpClient := http.Client{}
	graphqlClient := graphql.NewClient(baseUrl, &httpClient)

	return &Client{
		client: graphqlClient,
	}
}

func (c *Client) GetTransactions(ctx context.Context, filter domain.TransactionFilter, limit int, skip int) ([]domain.Transaction, error) {
	txs, err := GetTransactions(ctx, c.client, &filter, &limit, &skip)
	if err != nil {
		return nil, err
	}

	return txs.GetTransactions, nil
}
