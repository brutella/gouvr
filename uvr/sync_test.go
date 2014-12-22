package uvr

import (
	"github.com/stretchr/testify/assert"
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
	assert.False(t, syncDecoder.synced)
}

func TestValidSync(t *testing.T) {
	bitReceiver := NewTestBitReceiver()
	timeout := NewTimeout(244.0, 0.3)
	syncDecoder := NewSyncDecoder(bitReceiver, nil, timeout)
	signal := NewSignal(syncDecoder)

	writeWords([]big.Word{0, 0, 1, 1, 1, 1, 1, 1, 1, 1}, signal, timeout)
	assert.True(t, syncDecoder.synced)
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
		assert.Nil(t, err)
		bits = append(bits, bit)
	}
	bitReceiver := NewTestBitReceiver()
	timeout := NewTimeout(488.0, 0.3)
	syncDecoder := NewSyncDecoder(bitReceiver, nil, timeout)
	writeBits(bits, syncDecoder)
	assert.True(t, syncDecoder.synced)
}
