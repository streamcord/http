package http

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/streamcord/http/objects"
	"github.com/streamcord/http/ratelimit"
)

func (c *Client) AddRole(gID string, uID string, rID string) (*http.Response, error) {
	req := objects.Request{
		Endpoint: fmt.Sprintf("/guilds/%s/members/%s/roles/%s", gID, uID, rID),
		Method:   "PUT",
	}

	bID := fmt.Sprintf("/guilds/%s/members", gID)
	bucket := ratelimit.GetBucket(bID)
	if bucket != nil {
		if bucket.Remaining == 0 {
			wait := time.Duration(bucket.Reset - time.Now().Unix())
			// If wait is below 0 then that means it's already reset and we don't have to wait
			if wait > 0 {
				time.Sleep(wait * time.Second)
			}
		}
	}

	res, err := c.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		switch res.StatusCode {
		case 400:
			return nil, errors.New("malformed request payload")
		case 403:
			return nil, errors.New("missing required permission")
		default:
			return nil, fmt.Errorf("unknown error: %d", res.StatusCode)
		}
	}

	ratelimit.UpdateBucket(bID, res)
	return res, nil
}
