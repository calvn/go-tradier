package tradier

type UserService service

type User struct {
	Profile  *Profile  `json:"profile,omitempty"`
	Accounts *Accounts `json:"accounts,omitempty"`
}
