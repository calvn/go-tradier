package tradier

func (s *UserService) Orders() (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", "user/orders", nil)
	if err != nil {
		return nil, nil, err
	}

	o := &User{}

	resp, err := s.client.Do(req, o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, nil
}
