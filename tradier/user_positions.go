package tradier

import "time"

// FIXME: This needs to handle case where there is only 1 account entry
//        In that case, account is an object and not an array
type Positions struct {
	Accounts *struct {
		AccountEntries []PositionsAccountEntry `json:"account,omitempty"`
	} `json:"accounts,omitempty"`
}

type PositionsAccountEntry struct {
	AccountNumber *string `json:"account_number,omitempty"`
	Positions     *struct {
		PositionEntries []PositionEntry `json:"position,omitempty"`
	} `json:"positions,omitempty"`
}

type PositionEntry struct {
	CostBasis    *float64   `json:"cost_basis,omitempty"`
	DateAcquired *time.Time `json:"date_acquired,omitempty"`
	ID           *int       `json:"id,omitempty"`
	Quantity     *int       `json:"quantity,omitempty"`
	Symbol       *string    `json:"symbol,omitempty"`
}

func (s *UserService) Positions() (*Positions, *Response, error) {
	req, err := s.client.NewRequest("GET", "user/positions", nil)
	if err != nil {
		return nil, nil, err
	}

	p := &Positions{}

	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}
