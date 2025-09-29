package client

import "net/http"

type Client struct {
	httpClient *http.Client
	BaseURL    string
}

func NewClient(baseURL string) *Client {
	return &Client{
		httpClient: &http.Client{},
		BaseURL:    baseURL,
	}
}
