package steam

func (s *Steam) GetAppList(appList *AppList) error {
	if err := s.getRequest(APIEndpointGetAppList, nil, nil); err != nil {
		return err
	}

	return nil
}
