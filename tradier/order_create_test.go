package tradier

import (
	"net/http"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestOrder_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts/TEST/orders", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testFormValues(t, r, values{
			"class":    "equity",
			"symbol":   "AAPL",
			"duration": "day",
			"side":     "buy",
			"quantity": "1",
			"type":     "market",
		})

		w.WriteHeader(http.StatusOK)
		w.Write(marketsTimeSalesJSON)
	})

	params := &OrderParams{
		Class:    "equity",
		Symbol:   "AAPL",
		Duration: "day",
		Side:     "buy",
		Quantity: 1,
		Type:     "market",
	}

	got, _, err := client.Order.Create("TEST", params)
	if err != nil {
		t.Errorf("Markets.TimeSales returned error: %v", err)
	}

	if diff := pretty.Compare(wantOrdersCreate, got); diff != "" {
		t.Errorf("diff: %s", diff)
	}
}

var orderCreateJSON = []byte(`{
  "order": {
    "commission": 1,
    "cost": 122.95,
    "duration": "day",
    "extended_hours": false,
    "fees": 0,
    "margin_change": 121.95,
    "class": "equity",
    "order_cost": 121.95,
    "quantity": 1,
    "request_date": "2017-01-28T04:09:18.350Z",
    "result": true,
    "side": "buy",
    "status": "ok",
    "symbol": "AAPL",
    "type": "market"
  }
}`)

var wantOrdersCreate = &Order{
	Symbol: String("AAPL"),
}
