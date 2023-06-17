package steam

import (
	"fmt"
	"net/url"
)

type GetDOTA2MatchHistoryBySequenceNumResponse struct {
	Result struct {
		Status  int `json:"status"`
		Matches []struct {
			Players []struct {
				AccountID       int64 `json:"account_id"`
				PlayerSlot      int   `json:"player_slot"`
				TeamNumber      int   `json:"team_number"`
				TeamSlot        int   `json:"team_slot"`
				HeroID          int   `json:"hero_id"`
				Item0           int   `json:"item_0"`
				Item1           int   `json:"item_1"`
				Item2           int   `json:"item_2"`
				Item3           int   `json:"item_3"`
				Item4           int   `json:"item_4"`
				Item5           int   `json:"item_5"`
				Backpack0       int   `json:"backpack_0"`
				Backpack1       int   `json:"backpack_1"`
				Backpack2       int   `json:"backpack_2"`
				ItemNeutral     int   `json:"item_neutral"`
				Kills           int   `json:"kills"`
				Deaths          int   `json:"deaths"`
				Assists         int   `json:"assists"`
				LeaverStatus    int   `json:"leaver_status"`
				LastHits        int   `json:"last_hits"`
				Denies          int   `json:"denies"`
				GoldPerMin      int   `json:"gold_per_min"`
				XpPerMin        int   `json:"xp_per_min"`
				Level           int   `json:"level"`
				NetWorth        int   `json:"net_worth"`
				AghanimsScepter int   `json:"aghanims_scepter"`
				AghanimsShard   int   `json:"aghanims_shard"`
				Moonshard       int   `json:"moonshard"`
			} `json:"players"`
			RadiantWin            bool `json:"radiant_win"`
			Duration              int  `json:"duration"`
			PreGameDuration       int  `json:"pre_game_duration"`
			StartTime             int  `json:"start_time"`
			MatchID               int  `json:"match_id"`
			MatchSeqNum           int  `json:"match_seq_num"`
			TowerStatusRadiant    int  `json:"tower_status_radiant"`
			TowerStatusDire       int  `json:"tower_status_dire"`
			BarracksStatusRadiant int  `json:"barracks_status_radiant"`
			BarracksStatusDire    int  `json:"barracks_status_dire"`
			Cluster               int  `json:"cluster"`
			FirstBloodTime        int  `json:"first_blood_time"`
			LobbyType             int  `json:"lobby_type"`
			HumanPlayers          int  `json:"human_players"`
			Leagueid              int  `json:"leagueid"`
			PositiveVotes         int  `json:"positive_votes"`
			NegativeVotes         int  `json:"negative_votes"`
			GameMode              int  `json:"game_mode"`
			Flags                 int  `json:"flags"`
			Engine                int  `json:"engine"`
			RadiantScore          int  `json:"radiant_score"`
			DireScore             int  `json:"dire_score"`
		} `json:"matches"`
	} `json:"result"`
}

func (s *Steam) GetDOTA2MatchHistoryBySequenceNum(startAtMatchSeqNum int64) (*GetDOTA2MatchHistoryBySequenceNumResponse, error) {
	var response GetDOTA2MatchHistoryBySequenceNumResponse
	query := url.Values{}
	query.Add("start_at_match_seq_num", fmt.Sprintf("%d", startAtMatchSeqNum))
	err := s.getRequest("IDOTA2Match_570/GetMatchHistoryBySequenceNum/v1", query, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
