package tradier

func (s *UserService) History() (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", "user/history", nil)
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
