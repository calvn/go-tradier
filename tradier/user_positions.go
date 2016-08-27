package tradier

func (s *UserService) Positions() (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", "user/positions", nil)
	if err != nil {
		return nil, nil, err
	}

	p := &User{}

	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}
