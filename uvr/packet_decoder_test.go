package uvr

import (
	"math/big"
	"testing"
)

func TestPacketDecoder(t *testing.T) {
	packetReceiver := &packetReceiver{}
	packetDecoder := NewPacketDecoder(packetReceiver, 2)
	bitsTimeout := NewTimeout(244.0, 0.3)
	byteDecoder := NewByteDecoder(packetDecoder, bitsTimeout)

	syncBitsTimeout := NewTimeout(488.0, 0.5)
	syncDecoder := NewSyncDecoder(byteDecoder, byteDecoder, syncBitsTimeout)

	signal := NewSignal(syncDecoder)

	writeWords([]big.Word{1, 1, 1, 1, 1, 1, 1, 1}, signal, syncBitsTimeout) // sync

	if is, want := syncDecoder.synced, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	writeWords([]big.Word{0, 1, 1, 1, 1, 1, 1, 1, 1, 1}, signal, bitsTimeout) // 0xFF
	writeWords([]big.Word{0, 1, 0, 1, 0, 1, 0, 1, 0, 1}, signal, bitsTimeout) // 0x0F

	if is, want := len(packetReceiver.packet), 2; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := packetReceiver.packet[0], Byte(0xFF); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := packetReceiver.packet[1], Byte(0x55); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
