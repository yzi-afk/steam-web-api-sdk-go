package steam

func (c *Client) GetAppList(appList *AppList) error {
	if err := c.getRequest(APIEndpointGetAppList, nil, nil); err != nil {
		return err
	}

	return nil
}
