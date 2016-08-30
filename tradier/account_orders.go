package tradier

import "fmt"

func (s *AccountService) Orders(id string) (*Orders, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	orders := &Orders{}

	resp, err := s.client.Do(req, orders)
	if err != nil {
		return nil, resp, err
	}

	return orders, resp, nil
}
