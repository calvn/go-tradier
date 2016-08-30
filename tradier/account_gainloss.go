package tradier

import "fmt"

func (s *AccountService) GainLoss(id string) (*GainLoss, *Response, error) {
	u := fmt.Sprintf("accounts/%s/gainloss", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	gainloss := &GainLoss{}

	resp, err := s.client.Do(req, gainloss)
	if err != nil {
		return nil, resp, err
	}

	return gainloss, resp, nil
}
