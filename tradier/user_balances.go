package tradier

func (s *UserService) Balances() (*Balances, *Response, error) {
	req, err := s.client.NewRequest("GET", "user/balances", nil)
	if err != nil {
		return nil, nil, err
	}

	b := &Balances{}

	resp, err := s.client.Do(req, b)
	if err != nil {
		return nil, resp, err
	}

	return b, resp, nil
}
