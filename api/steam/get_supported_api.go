package steam

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
	if err := c.getRequest(APIEndpointGetSupportedAPIList, nil, r); err != nil {
		return err
	}

	return nil
}
