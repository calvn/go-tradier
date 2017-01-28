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
		w.Write(orderCreateJSON)
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

func TestOrder_Create_preview(t *testing.T) {
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
			"preview":  "true",
		})

		w.WriteHeader(http.StatusOK)
		w.Write(orderCreatePreviewJSON)
	})

	params := &OrderParams{
		Class:    "equity",
		Symbol:   "AAPL",
		Duration: "day",
		Side:     "buy",
		Quantity: 1,
		Type:     "market",
		Preview:  true,
	}

	got, _, err := client.Order.Create("TEST", params)
	if err != nil {
		t.Errorf("Markets.TimeSales returned error: %v", err)
	}

	if diff := pretty.Compare(wantOrdersCreatePreview, got); diff != "" {
		t.Errorf("diff: %s", diff)
	}
}

var orderCreateJSON = []byte(`{
  "order": {
    "id": 250501,
    "partner_id": "8162b718-9b19-4905-aa1d-f8087590a079",
    "status": "ok"
  }
}`)

var wantOrdersCreate = &Order{
	ID:        Int(250501),
	PartnerID: String("8162b718-9b19-4905-aa1d-f8087590a079"),
	Status:    String("ok"),
}

var orderCreatePreviewJSON = []byte(`{
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

var wantOrdersCreatePreview = &Order{
	Class:         String("equity"),
	Commission:    Float64(1),
	Cost:          Float64(122.95),
	Duration:      String("day"),
	ExtendedHours: Bool(false),
	Fees:          Float64(0),
	MarginChange:  Float64(121.95),
	OrderCost:     Float64(121.95),
	Quantity:      Float64(1),
	Result:        Bool(true),
	Side:          String("buy"),
	Status:        String("ok"),
	Symbol:        String("AAPL"),
	Type:          String("market"),
}
