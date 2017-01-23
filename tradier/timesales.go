package tradier

import (
	"encoding/json"
	"time"
)

// Series represents the series JSON object from /markets/timesales.
type Series struct {
	Data []*Data `json:"data,omitempty"`
}

type series Series

// Data represents the data JSON object from /markets/timesales.
type Data struct {
	Price     *float64   `json:"price,omitempty"`
	Time      *time.Time `json:"time,omitempty"`
	Timestamp *int       `json:"timestamp,omitempty"`
	Volume    *int       `json:"volume,omitempty"`
}

type data Data

// UnmarshalJSON unmarshals series into Series object.
func (s *Series) UnmarshalJSON(b []byte) (err error) {
	seriesNull := ""
	seriesObj := series{}
	dataObj := data{}
	dataCol := []*Data{}

	// If series is a string, i.e. "null"
	if err = json.Unmarshal(b, &seriesNull); err == nil {
		return nil
	}

	// If series is non-empty
	if err = json.Unmarshal(b, &seriesObj); err == nil {
		*s = Series(seriesObj)
		return nil
	}

	// If data is an object
	if err = json.Unmarshal(b, &dataObj); err == nil {
		d := Data(dataObj)
		*s = Series{
			Data: []*Data{&d},
		}
		return nil
	}

	if err = json.Unmarshal(b, &dataCol); err == nil {
		*s = Series{
			Data: dataCol,
		}
	}

	return nil
}

// MarshalJSON marshals Series into its JSON representation.
func (s *Series) MarshalJSON() ([]byte, error) {
	// If []Event is empty
	if len(s.Data) == 0 {
		return json.Marshal("null")
	}

	// If []Event is size 1, return first and only object
	if len(s.Data) == 1 {
		return json.Marshal(map[string]interface{}{
			"data": s.Data[0],
		})
	}

	// Otherwise mashal entire History object normally
	return json.Marshal(*s)
}
