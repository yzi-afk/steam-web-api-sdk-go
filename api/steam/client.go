package steam

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
)

const (
	SteamAPIBaseURL = "https://api.steampowered.com/"

	APIEndpointGetAppList = "ISteamApps/GetAppList/v2/"
)

type Client struct {
	apiKey          string
	httpClient      *http.Client
	endpointBaseURL *url.URL
}

func New(apiKey string, httpClient *http.Client) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("apiKey is required")
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	u, err := url.Parse(SteamAPIBaseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		apiKey:          apiKey,
		httpClient:      httpClient,
		endpointBaseURL: u,
	}, nil
}

func (c *Client) endpointURL(endpoint string) *url.URL {
	u := *c.endpointBaseURL
	u.Path += endpoint
	return &u
}

func (c *Client) doRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	return c.httpClient.Do(req)
}

func (c *Client) getRequest(ctx context.Context, endpoint string, query url.Values) (*http.Response, error) {
	u := c.endpointURL(endpoint)
	if query != nil {
		u.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}

func (c *Client) postRequest(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	u := c.endpointURL(endpoint)

	req, err := http.NewRequest(http.MethodPost, u.String(), body)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}