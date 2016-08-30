package tradier

import "fmt"

func (s *AccountService) Positions(id string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/positions", id)
	return s.AccountRequest(u)
}
