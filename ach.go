// Code generated by riza; DO NOT EDIT.

package bnkdev

import (
	"context"
	"fmt"
)

// Each account can have multiple account and routing numbers. We recommend
// that you use a set per vendor. This is similar to how you use different
// passwords for different websites.
type ACHRoute struct {
	ID string `json:"id,omitempty"` // The ACH route's identifier.

	// The identifier for the account this ACH route belongs to.
	AccountID string `json:"account_id,omitempty"`

	// The path that can be used in the API or your dashboard for the ACH
	// route.
	Path string `json:"path,omitempty"`

	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,omitempty"`

	// This indicates if payments can be made with the ACH route. The possible
	// values are active, disabled, and canceled
	Status string `json:"status,omitempty"`

	AccountNumber string `json:"account_number,omitempty"` // The account number.
	Name          string `json:"name,omitempty"`           // The name you chose for the ACH route.
}

type CreateACHRouteRequest struct {
	AccountID string `json:"account_id,omitempty"`
	Name      string `json:"name,omitempty"` // The name you choose for the ACH route.
}

// Returns CreateACHRouteRequest with AccountID set to the empty string so that it's
// not included in the JSON request body.
func (r *CreateACHRouteRequest) body() interface{} {
	if r == nil {
		return r
	}
	req := *r
	req.AccountID = ""
	return &req
}

type ListACHRouteRequest struct {
	AccountID string `json:"account_id,omitempty"`
}

type ListACHRouteResponse struct {
	Data             []ACHRoute       `json:"data,omitempty"`
	ResponseMetadata ResponseMetadata `json:"response_metadata,omitempty"`
}

func (c *Client) CreateACHRoute(ctx context.Context, r *CreateACHRouteRequest) (*ACHRoute, error) {
	var resp ACHRoute
	if err := c.request(ctx, "POST", fmt.Sprintf("/accounts/%s/routes/achs", r.AccountID), nil, r.body(), &resp); err != nil {
		return nil, fmt.Errorf("CreateACHRoute: %w", err)
	}
	return &resp, nil
}

// Returns a list of agent objects that match the provided query.
func (c *Client) ListACHRoute(ctx context.Context, r *ListACHRouteRequest) (*ListACHRouteResponse, error) {
	var resp ListACHRouteResponse
	if err := c.request(ctx, "GET", fmt.Sprintf("/accounts/%s/routes/achs", r.AccountID), nil, nil, &resp); err != nil {
		return nil, fmt.Errorf("ListACHRoute: %w", err)
	}
	return &resp, nil
}