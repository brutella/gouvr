package uvr1611

import (
	"github.com/brutella/gouvr/uvr"
	"testing"
)

func TestPacket(t *testing.T) {
	bytes := []uvr.Byte{
		uvr.Byte(uvr.DeviceTypeUVR1611),
		uvr.Byte(0x00),
		uvr.Byte(0x00),
		uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00),
		uvr.Byte(0xFA), uvr.Byte(0x20),
		uvr.Byte(0xAF), uvr.Byte(0x20),
		uvr.Byte(0x11), uvr.Byte(0x20),
		uvr.Byte(0x22), uvr.Byte(0x20),
		uvr.Byte(0x33), uvr.Byte(0x20),
		uvr.Byte(0x44), uvr.Byte(0x20),
		uvr.Byte(0x55), uvr.Byte(0x20),
		uvr.Byte(0x00), uvr.Byte(0x00),
		uvr.Byte(0x00), uvr.Byte(0x00),
		uvr.Byte(0x00), uvr.Byte(0x00),
		uvr.Byte(0x00), uvr.Byte(0x00),
		uvr.Byte(0x00), uvr.Byte(0x00),
		uvr.Byte(0x00), uvr.Byte(0x00),
		uvr.Byte(0x00), uvr.Byte(0x00),
		uvr.Byte(0x00), uvr.Byte(0x00),
		uvr.Byte(0x00), uvr.Byte(0x00),
		uvr.Byte(0x55), uvr.Byte(0x55), // 0101 0101 0101 0101
		uvr.Byte(0x00),
		uvr.Byte(0x00),
		uvr.Byte(0x00),
		uvr.Byte(0x00),
		uvr.Byte(0x00),
		uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00),
		uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00),
	}
	bytes = append(bytes, ChecksumFromBytes(bytes))

	packet, err := PacketFromBytes(bytes)
	if err != nil {
		t.Fatal(err)
	}

	if is, want := packet.DeviceType, uvr.DeviceTypeUVR1611; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := uvr.UInt16FromValue(packet.Input1), uint16(0x20FA); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := uvr.UInt16FromValue(packet.Input2), uint16(0x20AF); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := uvr.UInt16FromValue(packet.Input3), uint16(0x2011); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := uvr.UInt16FromValue(packet.Input4), uint16(0x2022); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := uvr.UInt16FromValue(packet.Input5), uint16(0x2033); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := uvr.UInt16FromValue(packet.Input6), uint16(0x2044); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := uvr.UInt16FromValue(packet.Input7), uint16(0x2055); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
