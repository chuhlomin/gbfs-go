package gbfs

import (
	"strings"
	"time"
)

type Clock time.Time

const clockFormat = "15:04:05"

func (clock *Clock) UnmarshalJSON(input []byte) error {
	newTime, err := time.Parse(clockFormat, strings.Trim(string(input), `"`))
	if err != nil {
		return err
	}

	*(*time.Time)(clock) = newTime
	return nil
}

func (t Clock) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t *Clock) String() string {
	return "\"" + time.Time(*t).Format(clockFormat) + "\""
}
