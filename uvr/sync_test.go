package uvr

import (
	"math/big"
	"strings"
	"testing"
)

func TestInvalidSync(t *testing.T) {
	bitReceiver := NewTestBitReceiver()
	timeout := NewTimeout(244.0, 0.3)
	syncDecoder := NewSyncDecoder(bitReceiver, nil, timeout)
	signal := NewSignal(syncDecoder)

	writeWords([]big.Word{1, 1, 1, 0, 1, 1, 1, 1}, signal, timeout)

	if is, want := syncDecoder.synced, false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestValidSync(t *testing.T) {
	bitReceiver := NewTestBitReceiver()
	timeout := NewTimeout(244.0, 0.3)
	syncDecoder := NewSyncDecoder(bitReceiver, nil, timeout)
	signal := NewSignal(syncDecoder)

	writeWords([]big.Word{0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, signal, timeout)

	if is, want := syncDecoder.synced, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestSyncFromLog(t *testing.T) {
	str := `1409826232854399456 1
1409826232856732000 1
1409826232859170430 1
1409826232861549101 1
1409826232863999654 1
1409826232866443723 1
1409826232868860208 1
1409826232871303484 1`

	lines := strings.Split(str, "\n")
	bits := make([]Bit, len(lines))
	for _, line := range lines {
		bit, err := BitFromLogString(line)

		if err != nil {
			t.Fatal(err)
		}

		bits = append(bits, bit)
	}
	bitReceiver := NewTestBitReceiver()
	timeout := NewTimeout(488.0, 0.3)
	syncDecoder := NewSyncDecoder(bitReceiver, nil, timeout)
	writeBits(bits, syncDecoder)

	if is, want := syncDecoder.synced, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
