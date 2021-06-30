package objects

type EmbedAuthor struct {
	IconURL string `json:"icon_url"`
	Name    string `json:"name"`
	URL     string `json:"url"`
}

type EmbedField struct {
	Inline bool   `json:"inline"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

type EmbedFooter struct {
	IconURL string `json:"icon_url"`
	Text    string `json:"text"`
}

type EmbedMedia struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type EmbedProvider struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Embed struct {
	Author      EmbedAuthor   `json:"author"`
	Color       int           `json:"color"`
	Description string        `json:"description"`
	Fields      []EmbedField  `json:"fields"`
	Footer      EmbedFooter   `json:"footer"`
	Image       EmbedMedia    `json:"image"`
	Provider    EmbedProvider `json:"provider"`
	Thumbnail   EmbedMedia    `json:"thumbnail"`
	Timestamp   string        `json:"timestamp"`
	Title       string        `json:"title"`
	Type        string        `json:"type"`
	URL         string        `json:"url"`
	Video       EmbedMedia    `json:"video"`
}
