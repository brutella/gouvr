package uvr

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	tm := time.Now()
	ntm := NewTimeForUnixNano(time.Duration(tm.UnixNano()))

	if is, want := milliseconds(ntm), milliseconds(tm); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
