package tradier

import "time"

type History struct {
	Event []Event `json:"event,omitempty"`
}

// FIXME: Handle array/object switch case
type Event struct {
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
