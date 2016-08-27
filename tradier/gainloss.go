package tradier

import "time"

type GainLoss struct {
	ClosedPosition []ClosedPosition `json:"closed_position,omitempty"`
}

// FIXME: Handle array/object switch case
type ClosedPosition struct {
	ClosedDate      *time.Time `json:"close_date,omitempty"`
	Cost            *float64   `json:"cost,omitempty"`
	GainLoss        *float64   `json:"gain_loss,omitempty"`
	GainLossPercent *float64   `json:"gain_loss_percent,omitempty"`
	OpenDate        *time.Time `json:"open_date,omitempty"`
	Proceeds        *float64   `json:"proceeds,omitempty"`
	Quantity        *int       `json:"quantity,omitempty"`
	Symbol          *string    `json:"symbol,omitempty"`
	Term            *int       `json:"term,omitempty"`
}
