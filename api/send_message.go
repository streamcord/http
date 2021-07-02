package http

import (
	"errors"
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
