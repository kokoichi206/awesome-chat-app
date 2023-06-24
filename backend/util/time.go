package util

import "time"

func FromISO8601(s string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, s)
}

func ToISO8601(t time.Time) string {
	return t.Format(time.RFC3339Nano)
}
