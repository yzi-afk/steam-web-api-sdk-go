package steam

import (
	"fmt"
)

type ServerInfoResponse struct {
	ServerTime       int64  `json:"servertime"`
	ServerTimeString string `json:"servertimestring"`
}

func (c *Client) GetServerInfo(r *ServerInfoResponse) error {
	resp, err := c.httpClient.Get(fmt.Sprint(c.endpointURL(APIEndpointGetServerInfo), "?key=", c.apiKey))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ParseJson(resp.Body, r)

	return nil
}
