package bnkdev

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

type Client struct {
	Base            string
	HTTP            http.Client
	EnableTelemetry bool

	key     string
	metrics chan *timing
}

func NewClient(key string) *Client {
	c := &Client{
		Base:            "https://api.bnk.dev",
		EnableTelemetry: true,
		key:             key,
	}
	return c
}

type Error struct {
	message string `json:"message"`
	code    int    `json:"code"`
}

func (e Error) Error() string {
	return e.message
}

func (c *Client) request(ctx context.Context, method, path string, params, in interface{}, out interface{}) error {
	var body io.Reader
	if in != nil {
		payload, err := json.Marshal(in)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(payload)
	}
	if params != nil {
		v, err := query.Values(params)
		if err != nil {
			return err
		}
		path += "?" + v.Encode()
	}
	req, err := http.NewRequest(method, c.Base+path, body)
	if err != nil {
		return err
	}

	var cleanup func(*http.Response)
	if c.EnableTelemetry {
		req, cleanup = withTelemetry(ctx, c, req)
	}

	req.Header.Set("Authorization", "Bearer "+c.key)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HTTP.Do(req.WithContext(ctx))

	if c.EnableTelemetry {
		defer func() {
			cleanup(resp)
		}()
	}

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		message, _ := ioutil.ReadAll(resp.Body)
		return Error{message: string(message), code: resp.StatusCode}
	}
	if out == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(out)
}
