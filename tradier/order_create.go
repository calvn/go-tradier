package tradier

import (
	"fmt"
	"net/url"
	"strconv"
)

func (s *OrderService) Create(accountId, class, symbol, duration, side string, quantity int, orderType string, price, stop float64, optionSymbol string) (*Order, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders", accountId)

	// Populate data
	data := url.Values{}
	data.Set("class", class)
	data.Set("symbol", symbol)
	data.Set("duration", duration)
	data.Set("side", side)
	data.Set("quantity", strconv.Itoa(quantity))
	data.Set("type", orderType)

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
