package steam

import (
	"net/url"
)

type GetDOTA2MatchHistoryResponse struct {
	Result struct {
		Status           int `json:"status"`
		NumResults       int `json:"num_results"`
		TotalResults     int `json:"total_results"`
		ResultsRemaining int `json:"results_remaining"`
		Matches          []struct {
			MatchID       int64 `json:"match_id"`
			MatchSeqNum   int64 `json:"match_seq_num"`
			StartTime     int   `json:"start_time"`
			LobbyType     int   `json:"lobby_type"`
			RadiantTeamID int   `json:"radiant_team_id"`
			DireTeamID    int   `json:"dire_team_id"`
			Players       []struct {
				AccountID  int64 `json:"account_id,omitempty"`
				PlayerSlot int   `json:"player_slot"`
				TeamNumber int   `json:"team_number"`
				TeamSlot   int   `json:"team_slot"`
				HeroID     int   `json:"hero_id"`
			} `json:"players"`
		} `json:"matches"`
	} `json:"result"`
}

func (c *Client) GetDOTA2MatchHistory() (*GetDOTA2MatchHistoryResponse, error) {
	var response GetDOTA2MatchHistoryResponse
	err := c.getRequest("IDOTA2Match_570/GetMatchHistory/v1", url.Values{}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
