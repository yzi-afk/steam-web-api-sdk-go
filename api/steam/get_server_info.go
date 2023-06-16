package steam

type ServerInfoResponse struct {
	ServerTime       int64  `json:"servertime"`
	ServerTimeString string `json:"servertimestring"`
}

func (c *Client) GetServerInfo(r *ServerInfoResponse) error {
	if err := c.getRequest(APIEndpointGetServerInfo, nil, r); err != nil {
		return err
	}

	return nil
}
