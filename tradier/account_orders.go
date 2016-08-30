package tradier

import "fmt"

func (s *AccountService) Orders(id string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/orders", id)
	return s.AccountRequest(u)
}
