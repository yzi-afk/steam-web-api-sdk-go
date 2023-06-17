package steam

type ServerInfoResponse struct {
	ServerTime       int64  `json:"servertime"`
	ServerTimeString string `json:"servertimestring"`
}

func (s *Steam) GetServerInfo(r *ServerInfoResponse) error {
	if err := s.getRequest(APIEndpointGetServerInfo, nil, r); err != nil {
		return err
	}

	return nil
}
