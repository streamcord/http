package payloads

import (
	"github.com/streamcord/http/objects"
)

type SendMessage struct {
	AllowedMentions objects.AllowedMentions `json:"allowed_mentions"`
	Content         string                  `json:"content"`
	Embeds          []objects.Embed         `json:"embeds"`
}
