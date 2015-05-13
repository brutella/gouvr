package uvr

import (
	"math/big"
	"testing"
)

func TestInvalidBytesConsumption(t *testing.T) {
	byteReceiver := NewTestByteReceiver()
	timeout := NewTimeout(244.0, 0.3)
	byteDecoder := NewByteDecoder(byteReceiver, timeout)
	syncDecoder := NewSyncDecoder(byteDecoder, byteDecoder, timeout)
	signal := NewSignal(syncDecoder)

	writeWords([]big.Word{1, 1, 1, 1, 1, 1, 1, 1}, signal, timeout) // sync

	if is, want := syncDecoder.synced, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	writeWords([]big.Word{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, signal, timeout) // wrong start byte resets the sync

	if is, want := syncDecoder.synced, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
