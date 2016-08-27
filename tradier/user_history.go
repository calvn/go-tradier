package tradier

func (s *UserService) History() (*History, *Response, error) {
	req, err := s.client.NewRequest("GET", "user/history", nil)
	if err != nil {
		return nil, nil, err
	}

	h := &History{}

	resp, err := s.client.Do(req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, nil
}
