package tradier

import "fmt"

func (s *AccountService) Balances(id string) (*Balances, *Response, error) {
	u := fmt.Sprintf("accounts/%s/balances", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	balances := &Balances{}

	resp, err := s.client.Do(req, balances)
	if err != nil {
		return nil, resp, err
	}

	return balances, resp, nil
}
