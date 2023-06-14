package steam

import (
	"context"
	"encoding/json"
	"io"
)

func (c *Client) GetAppList(appList *AppList) error {
	ctx := context.Background()
	res, err := c.getRequest(ctx, APIEndpointGetAppList, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	bodyJson := struct {
		AppList AppList `json:"applist"`
	}{}

	err = json.Unmarshal(b, &bodyJson)
	if err != nil {
		return err
	}

	return nil
}
