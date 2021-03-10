package cmp

import "time"

func Date(a, b time.Time) bool {
	return a.Truncate(time.Hour*24).Unix() == b.Truncate(time.Hour*24).Unix()
}
