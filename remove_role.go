package http

import (
	"fmt"
	"net/http"

	"github.com/streamcord/http/objects"
)

func (c *Client) RemoveRole(gID string, uID string, rID string) (*http.Response, error) {
	req := objects.Request{
		Endpoint:        fmt.Sprintf("/guilds/%s/members/%s/roles/%s", gID, uID, rID),
		Method:          "DELETE",
		RatelimitBucket: fmt.Sprintf("/guilds/%s/members", gID),
	}

	res, err := c.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
