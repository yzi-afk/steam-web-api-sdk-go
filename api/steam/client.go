package steam

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	SteamAPIBaseURL = "https://api.steampowered.com/"

	APIEndpointGetAppList          = "ISteamApps/GetAppList/v2/"
	APIEndpointGetServerInfo       = "ISteamWebAPIUtil/GetServerInfo/v1/"
	APIEndpointGetSupportedAPIList = "ISteamWebAPIUtil/GetSupportedAPIList/v1/"
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

func (c *Client) getRequest(endpoint string, query url.Values, v interface{}) error {
	u := c.endpointURL(endpoint)

	q := u.Query()
	q.Set("key", c.apiKey)
	for k, v := range query {
		q.Set(k, v[0])
	}
	u.RawQuery = q.Encode()

	resp, err := c.httpClient.Get(u.String())
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	ParseJson(resp.Body, v)

	return nil
}
