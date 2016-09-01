package tradier

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// NOTE: Refer to https://godoc.org/github.com/google/go-querystring/query for building the struct mapping
type OrderParams struct {
	Class        string  `url:"class"`
	Symbol       string  `url:"symbol"`
	Duration     string  `url:"duration"`
	Side         string  `url:"side,omitempty"`
	Quantity     int     `url:"quantity,omitempty"`
	Type         string  `url:"type"`
	Price        float64 `url:"price,omitempty"`
	Stop         float64 `url:"stop,omitempty"`
	OptionSymbol string  `url:"option_symbol,omitempty"`

	//Specific to multileg orders
	MultiSide         []string `url:"side,omitempty,[]"`
	MultiQuantity     []int    `url:"quantity,omitempty,[]"`
	MultiOptionSymbol []string `url:"option_symbol,omitempty,[]"`
}

func (s *OrderService) Create(accountId string, params *OrderParams) (*Order, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders", accountId)

	// Populate data
	data, err := query.Values(params)
	if err != nil {
		return nil, nil, err
	}
	// data := url.Values{}
	// data.Set("class", class)
	// data.Set("symbol", symbol)
	// data.Set("duration", duration)
	// data.Set("side", side)
	// data.Set("quantity", strconv.Itoa(quantity))
	// data.Set("type", orderType)

	uri, err := url.Parse(u)
	if err != nil {
		return nil, nil, err
	}

	uri.RawQuery = data.Encode()

	req, err := s.client.NewRequest("POST", u, data.Encode())
	if err != nil {
		return nil, nil, err
	}

	o := &Order{}

	resp, err := s.client.Do(req, o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, nil
}

func (s *OrderService) CreateMarket() {

}

func (s *OrderService) CreateLimit() {

}
