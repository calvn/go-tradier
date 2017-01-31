package tradier

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

var timesalesJSONSingle = []byte(`{
	"series": {
		"data": {
      "price": 4.32,
      "time": "2017-01-27T09:30:00",
      "timestamp": 1485527400,
      "volume": 125
    }
	}
}`)

var timesalesJSONArray = []byte(`{
	"series": {
		"data": [
      {
        "price": 4.32,
        "time": "2017-01-27T09:30:00",
        "timestamp": 1485527400,
        "volume": 125
      },
      {
        "price": 4.34,
        "time": "2017-01-27T09:31:20",
        "timestamp": 1485527480,
        "volume": 100
      }
    ]
	}
}`)

var timesaleJSONEmpty = []byte(`{
	"series": "null"
}`)

var timesalesSingle = &Series{
	Data: []*Data{
		&Data{
			Price:     Float64(4.32),
			Time:      &Time{time.Date(2017, 01, 27, 9, 30, 0, 0, time.UTC)},
			Timestamp: Int(1485527400),
			Volume:    Int(125),
		},
	},
}

var timesalesArray = &Series{
	Data: []*Data{
		&Data{
			Price:     Float64(4.32),
			Time:      &Time{time.Date(2017, 01, 27, 9, 30, 0, 0, time.UTC)},
			Timestamp: Int(1485527400),
			Volume:    Int(125),
		},
		&Data{
			Price:     Float64(4.34),
			Time:      &Time{time.Date(2017, 01, 27, 9, 31, 20, 0, time.UTC)},
			Timestamp: Int(1485527480),
			Volume:    Int(100),
		},
	},
}

var timesalesEmpty = &Series{}

func TestTimesales_UnmarshalJSON_Single(t *testing.T) {
	want := timesalesSingle

	got := &Series{}
	err := json.Unmarshal(timesalesJSONSingle, got)
	if err != nil {
		t.Errorf("Series.UnmarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v want: %+v", got, want)
	}
}

func TestTimesales_MarshalJSON_Single(t *testing.T) {
	buf := &bytes.Buffer{}
	err := json.Compact(buf, timesalesJSONSingle)
	want := buf.Bytes()
	if err != nil {
		t.Error(err)
	}

	got, err := json.Marshal(timesalesSingle)
	if err != nil {
		t.Errorf("Series.MarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s want: %s", got, want)
	}
}

func TestTimesales_UnmarshalJSON_Array(t *testing.T) {
	want := timesalesArray

	got := &Series{}
	err := json.Unmarshal(timesalesJSONArray, got)
	if err != nil {
		t.Errorf("Series.UnmarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v want: %+v", got, want)
	}
}

func TestTimesales_MarshalJSON_Array(t *testing.T) {
	buf := &bytes.Buffer{}
	err := json.Compact(buf, timesalesJSONArray)
	want := buf.Bytes()
	if err != nil {
		t.Error(err)
	}

	got, err := json.Marshal(timesalesArray)
	if err != nil {
		t.Errorf("Series.MarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s want: %s", got, want)
	}
}

func TestTimesales_UnmarshalJSON_Null(t *testing.T) {
	want := timesalesEmpty

	got := &Series{}
	err := json.Unmarshal(timesaleJSONEmpty, got)
	if err != nil {
		t.Errorf("Series.UnmarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %+v want: %+v", got, want)
	}
}

func TestTimesales_MarshalJSON_Null(t *testing.T) {
	buf := &bytes.Buffer{}
	err := json.Compact(buf, timesaleJSONEmpty)
	want := buf.Bytes()
	if err != nil {
		t.Error(err)
	}

	got, err := json.Marshal(&timesalesEmpty)
	if err != nil {
		t.Errorf("Series.MarshalJSON error: %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %s want: %s", got, want)
	}
}
