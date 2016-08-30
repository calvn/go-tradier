package tradier

import "fmt"

func (s *AccountService) History(id string) (*History, *Response, error) {
	u := fmt.Sprintf("accounts/%s/history", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	history := &History{}

	resp, err := s.client.Do(req, history)
	if err != nil {
		return nil, resp, err
	}

	return history, resp, nil
}
