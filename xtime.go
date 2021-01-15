package xtime

import (
	"time"
)

var now func() time.Time = time.Now

// Mock overwrites return value of xtime.Now().
// You must not use this function except for in test.
func Mock(fn func() time.Time) {
	now = fn
}

// Now returns the value of time.Now().
// If the Mock function has been executed in advance, the value set by Mock is returned.
func Now() time.Time {
	return now()
}
