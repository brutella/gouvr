package uvr

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "math/big"
)

func TestInvalidBytesConsumption(t *testing.T) {    
    byteReceiver := NewTestByteReceiver()
    timeout     := NewTimeout(244.0, 0.3)
    byteDecoder     := NewByteDecoder(byteReceiver, timeout)
    syncDecoder     := NewSyncDecoder(byteDecoder, timeout)    
    signal := NewSignal(syncDecoder)
    
    writeWords([]big.Word{1,1,1,1,1,1,1,1}, signal, timeout) // sync
    assert.True(t, syncDecoder.synced)
    writeWords([]big.Word{1,1,1,1,1,1,1,1,1,1}, signal, timeout) // wrong start byte resets the sync
    assert.False(t, syncDecoder.synced)
}
