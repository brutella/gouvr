package uvr

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "math/big"
)

func TestPacketConsumer(t *testing.T) {    
    packetReceiver := &packetReceiver{}
    packetDecoder   := NewPacketDecoder(packetReceiver, 2)
    bitsTimeout     := NewTimeout(244.0, 0.3)
    byteDecoder     := NewByteDecoder(packetDecoder, bitsTimeout)
    
    syncBitsTimeout := NewTimeout(488.0, 0.5)
    syncDecoder     := NewSyncDecoder(byteDecoder, syncBitsTimeout)
    
    signal := NewSignal(syncDecoder)
    
    writeWords([]big.Word{1,1,1,1,1,1,1,1}, signal, syncBitsTimeout) // sync
    assert.True(t, syncDecoder.synced)
    
    writeWords([]big.Word{0,1,1,1,1,1,1,1,1,1}, signal, bitsTimeout) // 0xFF
    writeWords([]big.Word{0,1,0,1,0,1,0,1,0,1}, signal, bitsTimeout) // 0x0F
    
    assert.Equal(t, len(packetReceiver.packet), 2)
    assert.Equal(t, packetReceiver.packet[0], Byte(0xFF))
    assert.Equal(t, packetReceiver.packet[1], Byte(0x55))
}
