package tradier

func (s *UserService) GainLoss() (*GainLoss, *Response, error) {
	req, err := s.client.NewRequest("GET", "user/gainloss", nil)
	if err != nil {
		return nil, nil, err
	}

	gl := &GainLoss{}

	resp, err := s.client.Do(req, gl)
	if err != nil {
		return nil, resp, err
	}

	return gl, resp, nil
}
