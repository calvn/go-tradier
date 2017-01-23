package tradier

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

// TimeSalesParams specifies the query parameters for querying time and sales
// for a particular symbol.
type TimeSalesParams struct {
	Symbol        string `url:"symbol"`
	Interval      string `url:"interval,omitempty"`
	Start         string `url:"start,omitempty"`
	End           string `url:"end,omitempty"`
	SessionFilter string `url:"session_filter,omitempty"`
}

// TimeSales returns the time and sales for a given symbol.
func (s *MarketsService) TimeSales(params *TimeSalesParams) (*Series, *Response, error) {
	u := fmt.Sprintf("markets/timesales")

	// Populate data
	data, err := query.Values(params)
	if err != nil {
		return nil, nil, err
	}

	uri, err := url.Parse(u)
	if err != nil {
		return nil, nil, err
	}

	uri.RawQuery = data.Encode()

	req, err := s.client.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	series := &Series{}

	resp, err := s.client.Do(req, series)
	if err != nil {
		return nil, resp, err
	}

	return series, resp, nil
}
