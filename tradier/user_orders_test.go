package tradier

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestActivityService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/feeds", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(userOrdersJSON)
	})

	got, _, err := client.User.Orders()
	if err != nil {
		t.Errorf("User.Orders returned error: %v", err)
	}
	if want := wantUserOrders; !reflect.DeepEqual(got, want) {
		t.Errorf("User.Orders = %+v, want %+v", got, want.Accounts.Account)
	}
}

var userOrdersJSON = []byte(`{
  "accounts": {
    "account": [
      {
        "account_number": "6YA05991",
        "orders": {
          "order": {
            "id": 182042,
            "type": "market",
            "symbol": "GOOGL",
            "side": "buy",
            "quantity": 1,
            "status": "pending",
            "duration": "gtc",
            "avg_fill_price": 0,
            "exec_quantity": 0,
            "last_fill_price": 0,
            "last_fill_quantity": 0,
            "remaining_quantity": 1,
            "create_date": "2016-08-23T05:17:37.617Z",
            "transaction_date": "2016-08-23T12:15:07.268Z",
            "class": "equity"
          }
        }
      },
      {
        "account_number": "6YA00005",
        "orders": {
          "order": [
            {
              "id": 182043,
              "type": "market",
              "symbol": "GOOGL",
              "side": "buy_to_open",
              "quantity": 1,
              "status": "pending",
              "duration": "gtc",
              "avg_fill_price": 0,
              "exec_quantity": 0,
              "last_fill_price": 0,
              "last_fill_quantity": 0,
              "remaining_quantity": 1,
              "create_date": "2016-08-23T05:18:19.256Z",
              "transaction_date": "2016-08-23T12:15:07.784Z",
              "class": "option",
              "option_symbol": "GOOGL160826C00815000"
            }, {
              "id": 182381,
              "type": "market",
              "symbol": "AAPL",
              "side": "buy_to_open",
              "quantity": 1,
              "status": "pending",
              "duration": "gtc",
              "avg_fill_price": 0,
              "exec_quantity": 0,
              "last_fill_price": 0,
              "last_fill_quantity": 0,
              "remaining_quantity": 1,
              "create_date": "2016-08-23T19:06:38.430Z",
              "transaction_date": "2016-08-23T19:06:38.446Z",
              "class": "option",
              "option_symbol": "AAPL160826C00099500"
            }
          ]
        }
      },
      {
        "account_number": "6YA05708",
        "orders": "null"
      }
    ]
  }
}`)

var (
	createdDate    = time.Date(2016, 8, 23, 05, 17, 37, 617, time.UTC)
	transitionDate = time.Date(2016, 8, 25, 12, 15, 07, 268, time.UTC)
)

var wantUserOrders = &User{
	Accounts: &Accounts{
		Account: []Account{
			{
				AccountNumber: String("6YA05991"),
				Orders: &Orders{
					Order: []Order{
						{
							ID:                Int(182042),
							Type:              String("Market"),
							Symbol:            String("GOOGL"),
							Side:              String("buy"),
							Quantity:          Float64(1),
							Status:            String("pending"),
							Duration:          String("gtc"),
							AvgFillPrice:      Float64(0),
							ExecQuantity:      Float64(0),
							LastFillPrice:     Float64(0),
							LastFillQuantity:  Float64(0),
							RemainingQuantity: Float64(1),
							CreateDate:        &createdDate,
							TransactionDate:   &transitionDate,
							Class:             String("equity"),
						},
					},
				},
			},
		},
	},
}
