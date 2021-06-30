package http

import (
	"bytes"
	"errors"
	"net/http"
	"time"

	"github.com/streamcord/http/objects"

	json "github.com/json-iterator/go"
)

type Client struct {
	HTTPClient http.Client
	Token      string // This should also include the token's prefix, e.g. "Bot <token>"
	URL        string // This is the URL of the API
}

func (c *Client) MakeRequest(r objects.Request) (*http.Response, error) {
	body, err := json.Marshal(r.Payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(r.Method, c.URL+r.Endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set request headers - individual routes shouldn't need their own headers
	req.Header.Set("Authorization", c.Token)
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewClient(timeout time.Duration, token string) (*Client, error) {
	// If timeout is 0 then failed connections (e.g. an outage on Discord's end), will cause our connections to hang leading an abundance of problems.
	if timeout == 0 {
		return nil, errors.New("timeout cannot be 0 for safety reasons")
	}

	return &Client{
		HTTPClient: http.Client{
			Timeout: timeout,
		},
		Token: token,
		URL:   "https://discord.com/api/v9",
	}, nil
}
