package tradier

// FIXME: This needs to handle case where there is only 1 account entry
//        In that case, account is an object and not an array
type UserOrders struct {
	Accounts *struct {
		AccountEntries []OrdersAccountEntry `json:"account,omitempty"`
	} `json:"accounts,omitempty"`
}

// `orders` can be an object or string literal "null"
// `order` can be an object or array of objects
type OrdersAccountEntry struct {
	AccountNumber *string `json:"account_number,omitempty"`
	Orders        *Orders `json:"orders,omitempty"`
}

func (s *UserService) Orders() (*UserOrders, *Response, error) {
	req, err := s.client.NewRequest("GET", "user/orders", nil)
	if err != nil {
		return nil, nil, err
	}

	o := &UserOrders{}

	resp, err := s.client.Do(req, o)
	if err != nil {
		return nil, resp, err
	}

	return o, resp, nil
}
