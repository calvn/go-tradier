package tradier

import "time"

// FIXME: This needs to handle case where there is only 1 account entry
//        In that case, account is an object and not an array
type GainLoss struct {
	Accounts *struct {
		AccountEntries []PositionsAccountEntry `json:"account,omitempty"`
	} `json:"accounts,omitempty"`
}

type GainLossAccountEntry struct {
	AccountNumber *string `json:"account_number,omitempty"`
	GainLoss      *struct {
		ClosedPositions []ClosedPositionEntry `json:"closed_positions,omitempty"`
	} `json:"gainloss,omitempty"`
}

type ClosedPositionEntry struct {
	ClosedDate      *time.Time `json:"close_date,omitempty"`
	Cost            *float64   `json:"cost,omitempty"`
	GainLoss        *float64   `json:"gain_loss,omitempty"`
	GainLossPercent *float64   `json:"gain_loss_percent,omitempty"`
	OpenDate        *time.Time `json:"open_date,omitempty"`
	Proceeds        *float64   `json:"proceeds,omitempty"`
	Quantity        *int       `json:"quantity,omitempty"`
	Symbol          *string    `json:"symbol,omitempty"`
	Term            *int       `json:"term,omitempty"`
}

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
