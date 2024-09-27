package gqlgenc

import (
	"context"
	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/coretech/dtc-api-gateway/internal/app/gateway/domain"
	"github.com/coretech/dtc-api-gateway/pkg/client/hdm-go-transactions-api/gqlgenc/generated"
	"net/http"
)

//go:generate go run github.com/Yamashou/gqlgenc

type Client struct {
	client *generated.Client
}

func New(baseUrl string) *Client {
	c := clientv2.Client{
		Client:  http.DefaultClient,
		BaseURL: baseUrl,
	}

	client := generated.Client{
		Client: &c,
	}

	return &Client{
		client: &client,
	}
}

func (c *Client) GetTransactions(ctx context.Context, filter domain.TransactionFilter, limit int, skip int) ([]domain.Transaction, error) {
	return nil, nil
}
