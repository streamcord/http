package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/streamcord/http/objects"
)

func (c *Client) AddRole(gID string, uID string, rID string) (*http.Response, error) {
	req := objects.Request{
		Endpoint:        fmt.Sprintf("/guilds/%s/members/%s/roles/%s", gID, uID, rID),
		Method:          "PUT",
		RatelimitBucket: fmt.Sprintf("/guilds/%s/members", gID),
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

	return res, nil
}
