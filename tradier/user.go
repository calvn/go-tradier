package tradier

// UserService handles routes related to user inquiry
// from the Tradier API.
type UserService service

// User represents the `user` JSON object
type User struct {
	Profile  *Profile  `json:"profile,omitempty"`
	Accounts *Accounts `json:"accounts,omitempty"`
}

func (s *UserService) userRequest(uri string) (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", uri, nil)
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
