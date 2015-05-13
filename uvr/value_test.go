package uvr

import (
	"testing"
)

func TestValue(t *testing.T) {
	bytes := []Byte{Byte(0x11), Byte(0xFA)}
	value := NewValue(bytes)

	// 0xFA11 = 64017
	if is, want := UInt16FromValue(value), uint16(64017); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	// 0xFA11 = -31249 unsigned
	if is, want := Int16FromValue(value), int16(-((65536 / 2) - 31249)); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
