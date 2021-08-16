package http

import (
	"fmt"
	"net/http"

	"github.com/streamcord/http/objects"
	"github.com/streamcord/http/payloads"
)

func (c *Client) SendMessage(cID string, payload payloads.SendMessage) (*http.Response, error) {
	req := objects.Request{
		Endpoint:        fmt.Sprintf("/channels/%s/messages", cID),
		Method:          "POST",
		Payload:         payload,
		RatelimitBucket: fmt.Sprintf("/channels/%s/messages", cID),
	}

	res, err := c.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
