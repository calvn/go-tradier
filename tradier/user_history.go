package tradier

import "time"

// FIXME: This needs to handle case where there is only 1 account entry
//        In that case, account is an object and not an array
type History struct {
	Accounts *struct {
		AccountEntries []HistoryAccountEntry `json:"account,omitempty"`
	} `json:"accounts,omitempty"`
}

type HistoryAccountEntry struct {
	AccountNumber *string `json:"account_number,omitempty"`
	History       *struct {
		Event []HistoryEventEntry `json:"event,omitempty"`
	} `json:"history,omitempty"`
}

type HistoryEventEntry struct {
	Amount   *float64   `json:"amount,omitempty"`
	Date     *time.Time `json:"date,omitempty"`
	Interest *struct {
		Description *string `json:"description,omitempty"`
		Quantity    *int    `json:"quantity,omitempty"`
	} `json:"interest,omitempty"`
	Journal *struct {
		Quantity *int `json:"quantity,omitempty"`
	} `json:"journal,omitempty"`
	Option *struct {
		Description *string `json:"description,omitempty"`
		OptionType  *string `json:"option_type,omitempty"`
		Quantity    *int    `json:"quantity,omitempty"`
	} `json:"option,omitempty"`
	Type  *string `json:"type,omitempty"`
	Trade *struct {
		Commission  *float64 `json:"commission,omitempty"`
		Description *string  `json:"description,omitempty"`
		Price       *float64 `json:"price,omitempty"`
		Quantity    *int     `json:"quantity,omitempty"`
		Symbol      *string  `json:"symbol,omitempty"`
		TradeType   *string  `json:"trade_type,omitempty"`
	} `json:"trade,omitempty"`
}

func (s *UserService) History() (*History, *Response, error) {
	req, err := s.client.NewRequest("GET", "user/history", nil)
	if err != nil {
		return nil, nil, err
	}

	h := &History{}

	resp, err := s.client.Do(req, h)
	if err != nil {
		return nil, resp, err
	}

	return h, resp, nil
}
