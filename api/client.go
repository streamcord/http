package http

import (
	"bytes"
	"errors"
	"net/http"
	"time"

	"github.com/streamcord/http/objects"
	"github.com/streamcord/http/ratelimit"

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

	bucket := ratelimit.GetBucket(r.RatelimitBucket)
	if bucket != nil {
		if bucket.Remaining == 0 {
			wait := time.Duration(bucket.Reset - time.Now().Unix())
			// If wait is below 0 then that means it's already reset and we don't have to wait
			if wait > 0 {
				time.Sleep(wait * time.Second)
			}
		}
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	ratelimit.UpdateBucket(r.RatelimitBucket, res)
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
