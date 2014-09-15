package uvr

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "math/big"
)

func TestRecordAndReplayBitStream(t *testing.T) {
    bitReceiver := NewTestBitReceiver()
    filePath := RandomTempFilePath()
    logger := NewFileLogger(filePath, 255)
    handover := NewHandover(bitReceiver, logger)
    
    words := []big.Word{0,1,0,1,0}
    bits := make([]Bit, 0, len(words))
    for _, w := range words {
        bits = append(bits, NewBitFromWord(w))
    }
    writeBits(bits, handover)
    logger.Flush()
    assert.Equal(t, len(bitReceiver.bits), 5)
    
    // Read file
    bitReceiver = NewTestBitReceiver()
    replayer := NewReplayer(bitReceiver)
    assert.Nil(t, replayer.Replay(filePath))
    assert.Equal(t, len(bitReceiver.bits), 5)
}