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
	req.Header.Set("User-Agent", "Streamcord (https://github.com/streamcord/http, 1.0.0)")

	bucket := ratelimit.GetBucket(r.RatelimitBucket)
	if bucket != nil {
		if bucket.Remaining == 0 {
			wait := time.Duration(bucket.Reset - float64(time.Now().Unix()))
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

	// If we get a 500/502 error for some reason, we can try again a minute later
	// In this case, ratelimits shouldn't really be a problem for long-lasting problems since Discord likely won't be receiving the request
	// Therefore a minute delay should not cause any interference.
	if res.StatusCode == http.StatusInternalServerError || res.StatusCode == http.StatusBadGateway {
		time.Sleep(time.Minute)
		return c.MakeRequest(r)
	} else if res.StatusCode < http.StatusInternalServerError {
		// Update ratelimit state using the response headers.
		// We don't want to be calling this if we get a 5xx error since there won't be any ratelimit headers to handle.
		err := ratelimit.UpdateBucket(r.RatelimitBucket, res)
		if err != nil {
			return nil, err
		}
	}

	// If we've got here then we've hit a ratelimit. Oh dear.
	// So, we'll retry the request when we can.
	// We don't need to change anything if the ratelimit is global as the reset header will refer to the global ratelimit.
	if res.StatusCode == http.StatusTooManyRequests {
		bucket = ratelimit.GetBucket(r.RatelimitBucket)
		if bucket != nil {
			wait := time.Duration(bucket.Reset - float64(time.Now().Unix()))
			if wait > 0 {
				time.Sleep(wait * time.Second)
				return c.MakeRequest(r)
			}
		}
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
