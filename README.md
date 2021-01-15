# xtime

## Overview
This is very simple time library which provides mock of time.Now().

The package name is short so as not to burden the coding.

## Usage
You just need to replace the `time.Now()` part with `xtime.Now()`.

Test and demo code is below.
```go
package xtime_test

import (
	"testing"
	"time"
	"xtime"
)

func TestNow(t *testing.T) {
	// 'xtime.Now() == time.Now()' become false due to the μ second level execution time difference.
	if xtime.Now().Sub(time.Now()) > time.Second*1 {
		t.Error("Default xtime.Now() must be same to time.Now().")
	}
	
	mockTime := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	xtime.Mock(mockTime)
	if xtime.Now() != mockTime {
		t.Error("xtime.Now() must be same to MockTime.")
	}

}
```
