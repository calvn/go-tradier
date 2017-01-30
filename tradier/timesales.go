package tradier

import "encoding/json"

// Series represents the series JSON object from /markets/timesales.
type Series struct {
	Data []*Data `json:"data,omitempty"`
}

type series Series

// Data represents the data JSON object from /markets/timesales.
type Data struct {
	Price     *float64 `json:"price,omitempty"`
	Time      *Time    `json:"time,omitempty"`
	Timestamp *int     `json:"timestamp,omitempty"`
	Volume    *int     `json:"volume,omitempty"`
}

type data Data

// UnmarshalJSON unmarshals series into Series object.
func (s *Series) UnmarshalJSON(b []byte) error {
	var seriesCol struct {
		// S *Series `json:"series,omitempty"`
		S struct {
			D []*Data `json:"data,omitempty"`
		} `json:"series,omitempty"`
	}
	var seriesObj struct {
		S struct {
			D *Data `json:"data,omitempty"`
		} `json:"seris,omitempty"`
	}

	var seriesStr struct {
		S string `json:"series,omitempty"`
	}

	var err error

	if err = json.Unmarshal(b, &seriesCol); err == nil {
		*s = Series{
			Data: seriesCol.S.D,
		}
		return nil
	}

	if err = json.Unmarshal(b, &seriesObj); err == nil {
		*s = Series{
			Data: []*Data{
				seriesObj.S.D,
			},
		}
		return nil
	}

	if err = json.Unmarshal(b, &seriesStr); err == nil {
		return nil
	}

	return err
}

// MarshalJSON marshals Series into its JSON representation.
func (s *Series) MarshalJSON() ([]byte, error) {
	if len(s.Data) == 0 {
		return json.Marshal(map[string]interface{}{
			"series": "null",
		})
	}

	if len(s.Data) == 1 {
		return json.Marshal(map[string]interface{}{
			"series": map[string]interface{}{
				"data": *(s.Data)[0],
			},
		})
	}

	return json.Marshal(map[string]interface{}{
		"series": map[string]interface{}{
			"data": s.Data,
		},
	})
}
