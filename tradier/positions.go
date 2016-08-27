package tradier

import "time"

type Positions struct {
	Position []Position `json:"position,omitempty"`
}

// FIXME: Handle array/object case correctly
type Position struct {
	CostBasis    *float64   `json:"cost_basis,omitempty"`
	DateAcquired *time.Time `json:"date_acquired,omitempty"`
	ID           *int       `json:"id,omitempty"`
	Quantity     *int       `json:"quantity,omitempty"`
	Symbol       *string    `json:"symbol,omitempty"`
}
