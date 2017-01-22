package tradier

import "time"

// Series represents the series JSON object from /markets/timesales.
type Series struct {
	Data []*Data `json:"data,omitempty"`
}

// Data represents the data JSON object from /markets/timesales.
type Data struct {
	Price     *float64   `json:"price,omitempty"`
	Time      *time.Time `json:"time,omitempty"`
	Timestamp *int       `json:"timestamp,omitempty"`
	Volume    *int       `json:"volume,omitempty"`
}
