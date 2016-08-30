package tradier

import "fmt"

func (s *AccountService) History(id string) (*Account, *Response, error) {
	u := fmt.Sprintf("accounts/%s/history", id)
	return s.AccountRequest(u)
}
