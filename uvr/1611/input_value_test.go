package uvr1611

import (
	"github.com/brutella/gouvr/uvr"
	"testing"
)

func TestTemperatureInputValue(t *testing.T) {
	value := uvr.NewValue([]uvr.Byte{
		uvr.Byte(0x50), // 0101 0000
		uvr.Byte(0x20), // 0010 0000
	})

	input_type, f := DecodeInputValue(value)
	if is, want := input_type, InputTypeTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := f, float32(8.0); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestNegativeTemperatureInputValue(t *testing.T) {
	value := uvr.NewValue([]uvr.Byte{
		uvr.Byte(0xE6), // 1110 0110
		uvr.Byte(0xAF), // 1010 1111
	})

	input_type, f := DecodeInputValue(value)
	if is, want := input_type, InputTypeTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := f, float32(-2.6); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestRoomTemperatureInputValue(t *testing.T) {
	value := uvr.NewValue([]uvr.Byte{
		uvr.Byte(0x50), // 0101 0000
		uvr.Byte(0x74), // 0111 0100
	})
	input_type, f := DecodeInputValue(value)

	if is, want := input_type, InputTypeRoomTemperature; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := f, float32(8.0); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	mode := RoomTemperatureModeFromValue(value)

	if is, want := mode, RoomTemperatureModeLowering; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestDigitalInputValue(t *testing.T) {
	value := uvr.NewValue([]uvr.Byte{
		uvr.Byte(0x50), // 0101 0000
		uvr.Byte(0x90), // 1001 0000
	})

	if is, want := IsDigitalInputValueOn(value), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestUnusedInputValue(t *testing.T) {
	value := uvr.NewValue([]uvr.Byte{
		uvr.Byte(0x10), // 0001 0000
		uvr.Byte(0x00), // 0000 0000
	})

	if is, want := IsUnusedInputValue(value), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
