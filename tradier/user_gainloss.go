package tradier

func (s *UserService) GainLoss() (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", "user/gainloss", nil)
	if err != nil {
		return nil, nil, err
	}

	u := &User{}

	resp, err := s.client.Do(req, u)
	if err != nil {
		return nil, resp, err
	}

	return u, resp, nil
}
