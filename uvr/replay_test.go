package uvr

import (
	"math/big"
	"testing"
)

func TestRecordAndReplayBitStream(t *testing.T) {
	bitReceiver := NewTestBitReceiver()
	filePath := RandomTempFilePath()
	logger := NewFileLogger(filePath, 255)
	handover := NewHandover(bitReceiver, logger)

	words := []big.Word{0, 1, 0, 1, 0}
	bits := make([]Bit, 0, len(words))
	for _, w := range words {
		bits = append(bits, NewBitFromWord(w))
	}
	writeBits(bits, handover)
	logger.Flush()

	if is, want := len(bitReceiver.bits), 5; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	// Read file
	bitReceiver = NewTestBitReceiver()
	replayer := NewReplayer(bitReceiver)

	if x := replayer.Replay(filePath); x != nil {
		t.Fatal(x)
	}
	if is, want := len(bitReceiver.bits), 5; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
