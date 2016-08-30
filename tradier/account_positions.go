package tradier

import "fmt"

func (s *AccountService) Positions(id string) (*Positions, *Response, error) {
	u := fmt.Sprintf("accounts/%s/positions", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	positions := &Positions{}

	resp, err := s.client.Do(req, positions)
	if err != nil {
		return nil, resp, err
	}

	return positions, resp, nil
}
