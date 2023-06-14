package steam

import (
	"errors"
	"net/http"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func New(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("apiKey is required")
	}

	return &Client{
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}, nil
}
