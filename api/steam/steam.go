package steam

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	SteamAPIBaseURL = "https://api.steampowered.com/"

	APIEndpointGetAppList          = "ISteamApps/GetAppList/v2/"
	APIEndpointGetServerInfo       = "ISteamWebAPIUtil/GetServerInfo/v1/"
	APIEndpointGetSupportedAPIList = "ISteamWebAPIUtil/GetSupportedAPIList/v1/"

	APIEndpointDOTA2GetLiveLeagueGames      = "IDOTA2Match_570/GetLiveLeagueGames/v1/"
	APIEndpointGetDOTA2MatchDetails         = "IDOTA2Match_570/GetMatchDetails/v1/"
	APIEndpointGetDOTA2TeamInfoByTeamID     = "IDOTA2Match_570/GetTeamInfoByTeamID/v1"
	APIEndpointGetDOTA2MatchHistory         = "IDOTA2Match_570/GetMatchHistory/v1"
	APIEndpointGetDOTA2MatchHistoryBySeqNum = "IDOTA2Match_570/GetMatchHistoryBySequenceNum/v1"
)

type Steam struct {
	apiKey          string
	httpClient      *http.Client
	endpointBaseURL *url.URL
}

func New(apiKey string, httpClient *http.Client) (*Steam, error) {
	if apiKey == "" {
		return nil, errors.New("apiKey is required")
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	u, err := url.Parse(SteamAPIBaseURL)
	if err != nil {
		return nil, err
	}

	return &Steam{
		apiKey:          apiKey,
		httpClient:      httpClient,
		endpointBaseURL: u,
	}, nil
}

func (s *Steam) endpointURL(endpoint string) *url.URL {
	u := *s.endpointBaseURL
	u.Path += endpoint
	return &u
}

func (s *Steam) getRequest(endpoint string, query url.Values, v interface{}) error {
	u := s.endpointURL(endpoint)

	q := u.Query()
	q.Set("key", s.apiKey)
	for k, v := range query {
		q.Set(k, v[0])
	}
	u.RawQuery = q.Encode()

	resp, err := s.httpClient.Get(u.String())
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, v)
	if err != nil {
		return err
	}

	return nil
}

func (s *Steam) GetTeamInfoByTeamID(teamID uint64) (*DOTA2TeamInfo, error) {
	params := url.Values{}
	params.Add("start_at_team_id", fmt.Sprint(teamID))
	params.Add("teams_requested", "1")

	var response GetTeamInfoByTeamIDResponse
	if err := s.getRequest(APIEndpointGetDOTA2TeamInfoByTeamID, params, &response); err != nil {
		return nil, err
	}

	if len(response.Result.Teams) == 0 {
		return nil, errors.New("team not found")
	}

	return &response.Result.Teams[0], nil

}

func (s *Steam) GetSupportedAPIList(r *GetSupportedAPIListResponse) error {
	if err := s.getRequest(APIEndpointGetSupportedAPIList, nil, r); err != nil {
		return err
	}

	return nil
}

func (s *Steam) GetServerInfo(r *ServerInfoResponse) error {
	if err := s.getRequest(APIEndpointGetServerInfo, nil, r); err != nil {
		return err
	}

	return nil
}

func (s *Steam) GetLiveLeagueGames() (*GetLiveLeagueGamesResponse, error) {
	var result GetLiveLeagueGamesResponse
	if err := s.getRequest(APIEndpointDOTA2GetLiveLeagueGames, nil, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *Steam) GetDOTA2MatchHistory() (*GetDOTA2MatchHistoryResponse, error) {
	var response GetDOTA2MatchHistoryResponse
	err := s.getRequest(APIEndpointGetDOTA2MatchHistory, url.Values{}, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Steam) GetDOTA2MatchHistoryBySequenceNum(startAtMatchSeqNum int64) (*GetDOTA2MatchHistoryBySequenceNumResponse, error) {
	var response GetDOTA2MatchHistoryBySequenceNumResponse
	query := url.Values{}
	query.Add("start_at_match_seq_num", fmt.Sprintf("%d", startAtMatchSeqNum))
	err := s.getRequest(APIEndpointGetDOTA2MatchHistoryBySeqNum, query, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Steam) GetDOTA2MatchDetails(matchID uint64) (*GetDOTA2MatchDetailsResponse, error) {
	var result GetDOTA2MatchDetailsResponse

	query := url.Values{}
	query.Add("match_id", fmt.Sprintf("%d", matchID))

	if err := s.getRequest(APIEndpointGetDOTA2MatchDetails, query, &result); err != nil {
		return nil, err
	}

	return &result, nil

}

func (s *Steam) GetAppList(appList *AppList) error {
	if err := s.getRequest(APIEndpointGetAppList, nil, nil); err != nil {
		return err
	}

	return nil
}
