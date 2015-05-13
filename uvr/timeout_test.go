package uvr

import (
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	duration := time.Duration(2)
	deviation := 0.5 // 50%
	timeout := NewTimeout(duration, deviation)

	if is, want := timeout.min(), time.Duration(250*time.Millisecond); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := timeout.max(), time.Duration(750*time.Millisecond); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	now := time.Now()

	if is, want := timeout.IsFutureSince(now), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := timeout.IsPastSince(now), false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	time.Sleep(250 * time.Millisecond)

	if is, want := timeout.IsFutureSince(now), false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := timeout.IsPastSince(now), false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	time.Sleep(600 * time.Millisecond)

	if is, want := timeout.IsFutureSince(now), false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := timeout.IsPastSince(now), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
