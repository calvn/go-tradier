package tradier

import (
	"net/http"
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
)

func TestMarketsService_TimeSales(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/markets/timesales", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"symbol": "GNCA",
		})

		w.WriteHeader(http.StatusOK)
		w.Write(marketsTimeSalesJSON)
	})

	params := &TimeSalesParams{
		Symbol: "GNCA",
	}

	got, _, err := client.Markets.TimeSales(params)
	if err != nil {
		t.Errorf("Markets.TimeSales returned error: %v", err)
	}

	if diff := pretty.Compare(wantMarketsTimeSales, got); diff != "" {
		t.Errorf("diff: %s", diff)
	}
}

var marketsTimeSalesJSON = []byte(`{
  "series": {
    "data": [
      {
      "price": 4.2,
      "time": "2017-01-20T09:30:00",
      "timestamp": 1484922600,
      "volume": 1322
      },
      {
      "price": 4.2,
      "time": "2017-01-20T09:30:29",
      "timestamp": 1484922629,
      "volume": 100
      }
    ]
  }
}`)

var (
	data1Time = time.Date(2017, 1, 20, 9, 30, 0, 000000000, time.UTC)
	data2Time = time.Date(2017, 1, 20, 9, 30, 29, 000000000, time.UTC)
)

var wantMarketsTimeSales = &Series{
	Data: []*Data{
		{
			Price:     Float64(4.2),
			Time:      &data1Time,
			Timestamp: Int(1484922600),
			Volume:    Int(1322),
		},
		{
			Price:     Float64(4.2),
			Time:      &data2Time,
			Timestamp: Int(1484922629),
			Volume:    Int(100),
		},
	},
}
