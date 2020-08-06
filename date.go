package xlib

import "time"

// NowB - returns the beginning of today
func NowB() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}
