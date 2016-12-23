package tradier

import (
	"encoding/json"
	"time"
)

// Accounts holds a slice of *Account
type Accounts []*Account

// Account represents the account object
type Account struct {
	AccountNumber *string `json:"account_number,omitempty"`

	// Specific to orders
	Orders *Orders `json:"orders,omitempty"`

	// Specific to profile
	Classification *string    `json:"classification,omitempty"`
	DateCreated    *time.Time `json:"date_created,omitempty"`
	DayTrader      *bool      `json:"day_trader,omitempty"`
	OptionLevel    *int       `json:"option_level,omitempty"`
	Status         *string    `json:"status,omitempty"`
	Type           *string    `json:"type,omitempty"`
	LastUpateDate  *time.Time `json:"last_update_date,omitempty"`

	// Specific to positions
	Positions *Positions `json:"positions,omitempty"`

	// Specific to gainloss
	GainLoss *GainLoss `json:"gainloss,omitempty"`

	// Specific to history
	History *History `json:"history,omitempty"`

	// Specific to balances
	Balances  *Balances `json:"balances,omitempty"`
	unwrapped bool      // used internally
}

type account Account

// UnmarshalJSON unmarshals account into Account object
func (a *Account) UnmarshalJSON(b []byte) error {
	var ac struct {
		*account `json:"account,omitempty"`
	}
	aObj := account{}

	// If wrapped
	if err := json.Unmarshal(b, &ac); err == nil {
		// log.Println("here")
		// log.Println(ac)
		if ac.account != nil {
			*a = Account(*ac.account)
			return nil
		}
	}

	// If not wrapped in anything
	if err := json.Unmarshal(b, &aObj); err == nil {
		*a = Account(aObj)
		return nil
	}

	return nil
}

// MarshalJSON marshals Account into JSON
func (a *Account) MarshalJSON() ([]byte, error) {
	if a.unwrapped {
		return json.Marshal(*a)
	}

	return json.Marshal(map[string]interface{}{
		"account": *a,
	})
}

// UnmarshalJSON unmarshals accounts into Accounts object
func (a *Accounts) UnmarshalJSON(b []byte) (err error) {
	var ac struct {
		A struct {
			A []*Account `json:"account,omitempty"`
		} `json:"accounts,omitempty"`
	}
	var aObj struct {
		A struct {
			A *Account `json:"account,omitempty"`
		} `json:"accounts,omitempty"`
	}
	var aNull string

	var aUnwrapped struct {
		A []*Account `json:"account,omitempty"`
	}

	// If unwrapped from user object
	if err = json.Unmarshal(b, &aUnwrapped); err == nil {
		*a = aUnwrapped.A
		return nil
	}

	// If account is a string, i.e. "null"
	if err = json.Unmarshal(b, &aNull); err == nil {
		return nil
	}

	// If account is an array
	if err = json.Unmarshal(b, &ac); err == nil {
		*a = ac.A.A
		return nil
	}

	// If account is an object
	if err = json.Unmarshal(b, &aObj); err == nil {
		acc := make([]*Account, 0)
		acc = append(acc, aObj.A.A)
		*a = Accounts(acc)
		return nil
	}

	return nil
}

// MarshalJSON marshals Accounts into JSON
func (a *Accounts) MarshalJSON() ([]byte, error) {
	// Set wrapped to true to marshal differently
	for _, acc := range *a {
		acc.unwrapped = true
	}

	// If []Watchlist is empty
	if len(*a) == 0 {
		return json.Marshal(map[string]interface{}{
			"accounts": "null",
		})
	}

	// If []Watchlist is size 1, return first and only object
	if len(*a) == 1 {
		acc := *a
		return json.Marshal(map[string]interface{}{
			"accounts": map[string]interface{}{
				"account": acc[0],
			},
		})
	}

	// Otherwhise marshal normally
	return json.Marshal(map[string]interface{}{
		"account": *a,
	})
}
