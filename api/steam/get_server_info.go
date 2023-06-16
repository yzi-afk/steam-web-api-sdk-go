package steam

import (
	"encoding/json"
	"fmt"
	"io"
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

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, r)
	if err != nil {
		return err
	}

	return nil
}
