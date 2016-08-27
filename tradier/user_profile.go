package tradier

// FIXME: This needs to handle case where there is only 1 account entry
//        In that case, account is an object and not an array
// type Profile struct {
// 	Profile *struct {
// 		AccountEntries []ProfileAccountEntry `json:"account,omitempty"`
// 		ID             *string               `json:"id,omitempty"`
// 		Name           *string               `json:"name,omitempty"`
// 	} `json:"profile,omitempty"`
// }
//
// type ProfileAccountEntry struct {
// 	ccountNumber   *string    `json:"account_number,omitempty"`
// 	Classification *string    `json:"classification,omitempty"`
// 	DateCreated    *time.Time `json:"date_created,omitempty"`
// 	DayTrader      *bool      `json:"day_trader,omitempty"`
// 	OptionLevel    *int       `json:"option_level,omitempty"`
// 	Status         *string    `json:"status,omitempty"`
// 	Type           *string    `json:"type,omitempty"`
// 	LastUpateDate  *time.Time `json:"last_update_date,omitempty"`
// }

func (s *UserService) Profile() (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", "user/profile", nil)
	if err != nil {
		return nil, nil, err
	}

	p := &User{}

	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, nil
}
