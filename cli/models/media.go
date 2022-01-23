package models

type Media struct {
	Caption   string `json:"caption"`
	MediaType string `json:"media_type"`
	MediaUrl  string `json:"media_url"`
	Id        string `json:"id"`
	Timestamp string `json:"timestamp"`
}

type MediaResponse = []Media
