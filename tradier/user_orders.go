package tradier

// FIXME: This needs to handle case where there is only 1 account entry
//        In that case, account is an object and not an array
// type UserOrders struct {
// 	Accounts *Accounts `json:"accounts,omitempty"`
// }

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
