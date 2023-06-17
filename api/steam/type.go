package steam

type AppList struct {
	Apps []App `json:"apps"`
}

type App struct {
	AppID int    `json:"appid"`
	Name  string `json:"name"`
}

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

type ServerInfoResponse struct {
	ServerTime       int64  `json:"servertime"`
	ServerTimeString string `json:"servertimestring"`
}

type GetLiveLeagueGamesResponse struct {
	Result struct {
		Games []struct {
			Players []struct {
				AccountID int    `json:"account_id"`
				Name      string `json:"name"`
				HeroID    int    `json:"hero_id"`
				Team      int    `json:"team"`
			} `json:"players"`
			RadiantTeam struct {
				TeamName string `json:"team_name"`
				TeamID   int    `json:"team_id"`
				TeamLogo int64  `json:"team_logo"`
				Complete bool   `json:"complete"`
			} `json:"radiant_team,omitempty"`
			DireTeam struct {
				TeamName string `json:"team_name"`
				TeamID   int    `json:"team_id"`
				TeamLogo int64  `json:"team_logo"`
				Complete bool   `json:"complete"`
			} `json:"dire_team,omitempty"`
			LobbyID           int64 `json:"lobby_id"`
			MatchID           int64 `json:"match_id"`
			Spectators        int   `json:"spectators"`
			LeagueID          int   `json:"league_id"`
			LeagueNodeID      int   `json:"league_node_id"`
			StreamDelayS      int   `json:"stream_delay_s"`
			RadiantSeriesWins int   `json:"radiant_series_wins"`
			DireSeriesWins    int   `json:"dire_series_wins"`
			SeriesType        int   `json:"series_type"`
			Scoreboard        struct {
				Duration           float64 `json:"duration"`
				RoshanRespawnTimer int     `json:"roshan_respawn_timer"`
				Radiant            struct {
					Score         int `json:"score"`
					TowerState    int `json:"tower_state"`
					BarracksState int `json:"barracks_state"`
					Picks         []struct {
						HeroID int `json:"hero_id"`
					} `json:"picks"`
					Bans []struct {
						HeroID int `json:"hero_id"`
					} `json:"bans"`
					Players []struct {
						PlayerSlot       int     `json:"player_slot"`
						AccountID        int     `json:"account_id"`
						HeroID           int     `json:"hero_id"`
						Kills            int     `json:"kills"`
						Death            int     `json:"death"`
						Assists          int     `json:"assists"`
						LastHits         int     `json:"last_hits"`
						Denies           int     `json:"denies"`
						Gold             int     `json:"gold"`
						Level            int     `json:"level"`
						GoldPerMin       int     `json:"gold_per_min"`
						XpPerMin         int     `json:"xp_per_min"`
						UltimateState    int     `json:"ultimate_state"`
						UltimateCooldown int     `json:"ultimate_cooldown"`
						Item0            int     `json:"item0"`
						Item1            int     `json:"item1"`
						Item2            int     `json:"item2"`
						Item3            int     `json:"item3"`
						Item4            int     `json:"item4"`
						Item5            int     `json:"item5"`
						RespawnTimer     int     `json:"respawn_timer"`
						PositionX        float64 `json:"position_x"`
						PositionY        float64 `json:"position_y"`
						NetWorth         int     `json:"net_worth"`
					} `json:"players"`
					Abilities []struct {
						AbilityID    int `json:"ability_id"`
						AbilityLevel int `json:"ability_level"`
					} `json:"abilities"`
				} `json:"radiant"`
				Dire struct {
					Score         int `json:"score"`
					TowerState    int `json:"tower_state"`
					BarracksState int `json:"barracks_state"`
					Picks         []struct {
						HeroID int `json:"hero_id"`
					} `json:"picks"`
					Bans []struct {
						HeroID int `json:"hero_id"`
					} `json:"bans"`
					Players []struct {
						PlayerSlot       int     `json:"player_slot"`
						AccountID        int     `json:"account_id"`
						HeroID           int     `json:"hero_id"`
						Kills            int     `json:"kills"`
						Death            int     `json:"death"`
						Assists          int     `json:"assists"`
						LastHits         int     `json:"last_hits"`
						Denies           int     `json:"denies"`
						Gold             int     `json:"gold"`
						Level            int     `json:"level"`
						GoldPerMin       int     `json:"gold_per_min"`
						XpPerMin         int     `json:"xp_per_min"`
						UltimateState    int     `json:"ultimate_state"`
						UltimateCooldown int     `json:"ultimate_cooldown"`
						Item0            int     `json:"item0"`
						Item1            int     `json:"item1"`
						Item2            int     `json:"item2"`
						Item3            int     `json:"item3"`
						Item4            int     `json:"item4"`
						Item5            int     `json:"item5"`
						RespawnTimer     int     `json:"respawn_timer"`
						PositionX        float64 `json:"position_x"`
						PositionY        float64 `json:"position_y"`
						NetWorth         int     `json:"net_worth"`
					} `json:"players"`
					Abilities []struct {
						AbilityID    int `json:"ability_id"`
						AbilityLevel int `json:"ability_level"`
					} `json:"abilities"`
				} `json:"dire"`
			} `json:"scoreboard,omitempty"`
		} `json:"games"`
		Status int `json:"status"`
	} `json:"result"`
}

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

type GetDOTA2MatchDetailsResponse struct {
	Result struct {
		Players []struct {
			AccountID         int `json:"account_id"`
			PlayerSlot        int `json:"player_slot"`
			TeamNumber        int `json:"team_number"`
			TeamSlot          int `json:"team_slot"`
			HeroID            int `json:"hero_id"`
			Item0             int `json:"item_0"`
			Item1             int `json:"item_1"`
			Item2             int `json:"item_2"`
			Item3             int `json:"item_3"`
			Item4             int `json:"item_4"`
			Item5             int `json:"item_5"`
			Backpack0         int `json:"backpack_0"`
			Backpack1         int `json:"backpack_1"`
			Backpack2         int `json:"backpack_2"`
			ItemNeutral       int `json:"item_neutral"`
			Kills             int `json:"kills"`
			Deaths            int `json:"deaths"`
			Assists           int `json:"assists"`
			LeaverStatus      int `json:"leaver_status"`
			LastHits          int `json:"last_hits"`
			Denies            int `json:"denies"`
			GoldPerMin        int `json:"gold_per_min"`
			XpPerMin          int `json:"xp_per_min"`
			Level             int `json:"level"`
			NetWorth          int `json:"net_worth"`
			AghanimsScepter   int `json:"aghanims_scepter"`
			AghanimsShard     int `json:"aghanims_shard"`
			Moonshard         int `json:"moonshard"`
			HeroDamage        int `json:"hero_damage"`
			TowerDamage       int `json:"tower_damage"`
			HeroHealing       int `json:"hero_healing"`
			Gold              int `json:"gold"`
			GoldSpent         int `json:"gold_spent"`
			ScaledHeroDamage  int `json:"scaled_hero_damage"`
			ScaledTowerDamage int `json:"scaled_tower_damage"`
			ScaledHeroHealing int `json:"scaled_hero_healing"`
			AbilityUpgrades   []struct {
				Ability int `json:"ability"`
				Time    int `json:"time"`
				Level   int `json:"level"`
			} `json:"ability_upgrades"`
		} `json:"players"`
		RadiantWin            bool   `json:"radiant_win"`
		Duration              int    `json:"duration"`
		PreGameDuration       int    `json:"pre_game_duration"`
		StartTime             int    `json:"start_time"`
		MatchID               int64  `json:"match_id"`
		MatchSeqNum           int64  `json:"match_seq_num"`
		TowerStatusRadiant    int    `json:"tower_status_radiant"`
		TowerStatusDire       int    `json:"tower_status_dire"`
		BarracksStatusRadiant int    `json:"barracks_status_radiant"`
		BarracksStatusDire    int    `json:"barracks_status_dire"`
		Cluster               int    `json:"cluster"`
		FirstBloodTime        int    `json:"first_blood_time"`
		LobbyType             int    `json:"lobby_type"`
		HumanPlayers          int    `json:"human_players"`
		Leagueid              int    `json:"leagueid"`
		PositiveVotes         int    `json:"positive_votes"`
		NegativeVotes         int    `json:"negative_votes"`
		GameMode              int    `json:"game_mode"`
		Flags                 int    `json:"flags"`
		Engine                int    `json:"engine"`
		RadiantScore          int    `json:"radiant_score"`
		DireScore             int    `json:"dire_score"`
		RadiantTeamID         int    `json:"radiant_team_id"`
		RadiantName           string `json:"radiant_name"`
		RadiantLogo           int64  `json:"radiant_logo"`
		RadiantTeamComplete   int    `json:"radiant_team_complete"`
		DireTeamID            int    `json:"dire_team_id"`
		DireName              string `json:"dire_name"`
		DireLogo              int64  `json:"dire_logo"`
		DireTeamComplete      int    `json:"dire_team_complete"`
		RadiantCaptain        int    `json:"radiant_captain"`
		DireCaptain           int    `json:"dire_captain"`
		PicksBans             []struct {
			IsPick bool `json:"is_pick"`
			HeroID int  `json:"hero_id"`
			Team   int  `json:"team"`
			Order  int  `json:"order"`
		} `json:"picks_bans"`
	} `json:"result"`
}
