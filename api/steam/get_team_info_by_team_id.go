package steam

import (
	"fmt"
	"net/url"
)

type GetTeamInfoByTeamIDResponse struct {
	Result struct {
		Status int `json:"status"`
		Teams  []struct {
			Name             string `json:"name"`
			Tag              string `json:"tag"`
			Abbreviation     string `json:"abbreviation"`
			TimeCreated      int    `json:"time_created"`
			Logo             int64  `json:"logo"`
			LogoSponsor      int64  `json:"logo_sponsor"`
			CountryCode      string `json:"country_code"`
			URL              string `json:"url"`
			GamesPlayed      int    `json:"games_played"`
			Player0AccountID int    `json:"player_0_account_id,omitempty"`
			Player1AccountID int    `json:"player_1_account_id,omitempty"`
			Player2AccountID int    `json:"player_2_account_id,omitempty"`
			AdminAccountID   int    `json:"admin_account_id"`
			Player3AccountID int    `json:"player_3_account_id,omitempty"`
			Player4AccountID int    `json:"player_4_account_id,omitempty"`
			Player5AccountID int    `json:"player_5_account_id,omitempty"`
			Player6AccountID int    `json:"player_6_account_id,omitempty"`
		} `json:"teams"`
	} `json:"result"`
}

func (c *Client) GetTeamInfoByTeamID(teamID int) (*GetTeamInfoByTeamIDResponse, error) {
	params := url.Values{}
	params.Add("start_at_team_id", fmt.Sprint(teamID))
	params.Add("teams_requested", "1")

	var response GetTeamInfoByTeamIDResponse
	if err := c.getRequest("IDOTA2Match_570/GetTeamInfoByTeamID/v1", params, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
