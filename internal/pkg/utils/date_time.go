package utils

import "time"

func TimeToString(t time.Time) string {
	format := "Mon, 02 Jan 2006 03:04:05 PM"
	return t.Format(format)
}
