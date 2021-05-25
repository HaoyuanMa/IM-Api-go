package model

type Message struct {
	Type        string   `json:"type"`
	From        string   `json:"from"`
	To          []string `json:"to"`
	ContentType string   `json:"contentType"`
	Content     string   `json:"content"`
}
