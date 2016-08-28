package tradier

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

var accountsJSONSingle = []byte(`{
  "account": {
    "account_number": "6YA05708"
  }
}`)

var accountsJSONArray = []byte(`{
  "account": [{
    "account_number": "6YA05708"
  }, {
    "account_number": "6YA05709"
  }]
}`)

var accountsJSONNull = []byte(`"null"`)

var accountsSingle = &Accounts{
	Account: []Account{
		{
			AccountNumber: String("6YA05708"),
		},
	},
}

func TestAccounts_UnmarshalJSONSingle(t *testing.T) {
	want := accountsSingle

	got := &Accounts{}
	err := json.Unmarshal(accountsJSONSingle, got)
	if err != nil {
		t.Errorf("Accounts.UnmarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v want: %+v", got, want)
	}
}

func TestAccounts_MarshalJSONSingle(t *testing.T) {
	buf := &bytes.Buffer{}
	err := json.Compact(buf, accountsJSONSingle)
	want := buf.Bytes()
	if err != nil {
		t.Error(err)
	}

	got, err := json.Marshal(accountsSingle)
	if err != nil {
		t.Errorf("Accounts.MarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s want: %s", got, want)
	}
}

func TestAccounts_UnmarshalJSON_Array(t *testing.T) {
	want := &Accounts{
		Account: []Account{
			{
				AccountNumber: String("6YA05708"),
			},
			{
				AccountNumber: String("6YA05709"),
			},
		},
	}

	accountsJSON := []byte(`{
    "account": [{
      "account_number": "6YA05708"
    }, {
      "account_number": "6YA05709"
    }]
  }`)

	got := &Accounts{}
	err := json.Unmarshal(accountsJSON, got)
	if err != nil {
		t.Errorf("Accounts.UnmarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v want: %+v", got, want)
	}
}

func TestAccounts_UnmarshalJSON_Null(t *testing.T) {
	want := &Accounts{}

	accountsJSON := []byte(`"null"`)

	got := &Accounts{}
	err := json.Unmarshal(accountsJSON, got)
	if err != nil {
		t.Errorf("Accounts.UnmarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v want: %+v", got, want)
	}
}
