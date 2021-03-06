// Code generated by riza; DO NOT EDIT.

package bnkdev

import (
	"context"
	"encoding/json"
	"fmt"
)

type ListPendingTransactionsRequest struct {
	AccountID string `json:"account_id,omitempty"`
}

type ListTransactionsRequest struct {
	AccountID string `json:"account_id,omitempty"`
}

type ListTransactionsResponse struct {
	Data             []Transaction    `json:"data,omitempty"`
	ResponseMetadata ResponseMetadata `json:"response_metadata,omitempty"`
}

// Transactions are the immutable additions and removals of money from your
// bank account.
type Transaction struct {
	ID string `json:"id,omitempty"` // The account identifier.

	// The identifier for the account this transaction belongs to.
	AccountID string `json:"account_id,omitempty"`

	// The transaction amount in the minor unit of the account currency. For
	// dollars, for example, this is cents.
	Amount int `json:"amount,omitempty"`

	// For a transaction related to a transfer, this is the description you
	// provide. For a transaction related to a payment, this is the
	// description the vendor provides.
	Description string `json:"description,omitempty"`

	// The transaction path that can be used in the API or your dashboard.
	Path string `json:"path,omitempty"`

	// The identifier for the route this transaction came through. Routes are
	// things like cards and ACH details.
	RouteID string `json:"route_id,omitempty"`

	// This is an object giving more details on the network-level event that
	// caused the transaction. For example, for a card transaction this lists
	// the merchant's industry and location.
	Source json.RawMessage `json:"source,omitempty"`

	Date string `json:"date,omitempty"` // The ISO 8601 date on which the transaction occured.
}

// Returns a list of pending transations.
func (c *Client) ListPendingTransactions(ctx context.Context, r *ListPendingTransactionsRequest) (*ListTransactionsResponse, error) {
	var resp ListTransactionsResponse
	if err := c.request(ctx, "GET", fmt.Sprintf("/accounts/%s/pending-transactions", r.AccountID), nil, nil, &resp); err != nil {
		return nil, fmt.Errorf("ListPendingTransactions: %w", err)
	}
	return &resp, nil
}

// Returns a list of transactions.
func (c *Client) ListTransactions(ctx context.Context, r *ListTransactionsRequest) (*ListTransactionsResponse, error) {
	var resp ListTransactionsResponse
	if err := c.request(ctx, "GET", fmt.Sprintf("/accounts/%s/transactions", r.AccountID), nil, nil, &resp); err != nil {
		return nil, fmt.Errorf("ListTransactions: %w", err)
	}
	return &resp, nil
}
