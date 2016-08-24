package tradier

type Accounts struct {
	Account []Account `json:"account,omitempty"`
}

type Account struct {
	AccountNumber *string `json:"account_number,omitempty"`
	Orders        *Orders `json:"orders,omitempty"`
}
