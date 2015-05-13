package uvr

import (
	"testing"
)

func TestTimestampToString(t *testing.T) {
	bytes := []Byte{
		Byte(45), // min
		Byte(23), // h
		Byte(2),  // d
		Byte(12), // M
		Byte(14), // y
	}
	timestamp := NewTimestamp(bytes)

	if is, want := timestamp.ToString(), "2014-12-02 23:45:00"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
