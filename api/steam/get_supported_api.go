package steam

import (
	"fmt"
)

type GetSupportedAPIListResponse struct {
	Apilist struct {
		Interfaces []struct {
			Name    string `json:"name"`
			Methods []struct {
				Name       string `json:"name"`
				Version    int    `json:"version"`
				Httpmethod string `json:"httpmethod"`
				Parameters []struct {
					Name        string `json:"name"`
					Type        string `json:"type"`
					Optional    bool   `json:"optional"`
					Description string `json:"description"`
				} `json:"parameters"`
			} `json:"methods"`
		} `json:"interfaces"`
	} `json:"apilist"`
}

func (c *Client) GetSupportedAPIList(r *GetSupportedAPIListResponse) error {
	resp, err := c.httpClient.Get(fmt.Sprint(c.endpointURL(APIEndpointGetSupportedAPIList), "?key=", c.apiKey))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ParseJson(resp.Body, r)

	return nil
}
