package tradier

import (
	"net/http"
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
)

func TestActivityService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user/orders", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(userOrdersJSON)
	})

	got, _, err := client.User.Orders()
	if err != nil {
		t.Errorf("User.Orders returned error: %v", err)
	}

	if diff := pretty.Compare(wantUserOrders, got); diff != "" {
		t.Errorf("diff: %s", diff)
	}
}

var userOrdersJSON = []byte(`{
  "accounts": {
    "account": [
      {
        "account_number": "6YA05991",
        "orders": "null"
      },
      {
        "account_number": "6YA05708",
        "orders": "null"
      }
    ]
  }
}`)

var (
	createdDate    = time.Date(2016, 8, 23, 05, 17, 37, 617000000, time.UTC)
	transitionDate = time.Date(2016, 8, 23, 12, 15, 07, 268000000, time.UTC)
)

var wantUserOrders = &User{
	Accounts: &Accounts{
		{
			AccountNumber: String("6YA05991"),
			Orders:        &Orders{},
		},
		{
			AccountNumber: String("6YA05708"),
			Orders:        &Orders{},
		},
	},
}
