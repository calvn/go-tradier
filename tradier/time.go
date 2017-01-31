package tradier

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// Time is the custom time-formatted time
type Time struct {
	time.Time
}

func (t Time) String() string {
	return t.Time.String()
}

// UnmarshalJSON unmarshals custom JSON time format for timesales date
func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		*t = Time{time.Time{}}
		return nil
	}

	var parsedTime time.Time

	parsedTime, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}

	*t = Time{parsedTime}
	return nil
}

// MarshalJSON marshals Time into timesales compatible time format
func (t *Time) MarshalJSON() ([]byte, error) {
	if t != nil {
		time := t.Format("2006-01-02T15:04:05")
		log.Println(time)

		return json.Marshal(time)

	}

	// FIXME: This is probably wrong
	return []byte{}, fmt.Errorf("Invalid time")
}
