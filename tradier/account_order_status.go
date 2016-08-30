package tradier

import "fmt"

func (s *AccountService) OrderStatus(accountId, orderId string) (*Order, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders/%s", accountId, orderId)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	order := &Order{}

	resp, err := s.client.Do(req, order)
	if err != nil {
		return nil, resp, err
	}

	return order, resp, nil
}
