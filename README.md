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

type incrementalMock struct {
	i time.Duration
	t time.Time
}

func (m *incrementalMock) Now() time.Time {
	m.i++
	return m.t.Add(time.Second * m.i)
}

func TestNow(t *testing.T) {
	// case 1.
	// Default xtime.Now behavior
	// 'xtime.Now() == time.Now()' become false due to the μ second level execution time difference.
	if !(xtime.Now().Sub(time.Now()) <= time.Second*1) {
		t.Error("Default xtime.Now() must be same to time.Now().")
	}

	// case 2.
	// Constant value Mock
	mockTime := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	xtime.Mock(func() time.Time { return mockTime })
	if xtime.Now() != mockTime {
		t.Error("xtime.Now() must be same to MockTime.")
	}

	// case 3.
	// Incremental value Mock
	incMock := &incrementalMock{
		i: 0,
		t: mockTime,
	}
	xtime.Mock(incMock.Now)
	if xtime.Now().Sub(mockTime) != time.Second*1 {
		t.Error("xtime.Now() must be same to MockTime+1sec.")
	}
	if xtime.Now().Sub(mockTime) != time.Second*2 {
		t.Error("xtime.Now() must be same to MockTime+2sec.")
	}

	// case 4.
	// reset
	xtime.Mock(time.Now)
	// Same to 1st test case.
	if !(xtime.Now().Sub(time.Now()) <= time.Second*1) {
		t.Error("Default xtime.Now() must be same to time.Now().")
	}
}
```
