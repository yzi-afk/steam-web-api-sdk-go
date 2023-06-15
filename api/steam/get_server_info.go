package steam

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ServerInfoResponse struct {
	ServerTime       int64  `json:"servertime"`
	ServerTimeString string `json:"servertimestring"`
}

func (c *Client) GetServerInfo(r *ServerInfoResponse) error {
	resp, err := http.Get(fmt.Sprint(c.endpointURL(APIEndpointGetServerInfo), "?key=", c.apiKey))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	bodyJson := &ServerInfoResponse{}

	err = json.Unmarshal(b, &bodyJson)
	if err != nil {
		return err
	}

	return nil
}
