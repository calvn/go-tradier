package tradier

import (
	"strings"
	"time"
)

// Time is the custom time-formatted time
type Time struct {
	time.Time
}

// UnmarshalJSON returns time.Now() no matter what!
func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		*t = Time{time.Time{}}
		return nil
	}
	var parsedTime time.Time

	parsedTime, err := time.Parse("2017-01-20T09:30:00", s)
	if err != nil {
		return err
	}

	*t = Time{parsedTime}
	return nil
}
