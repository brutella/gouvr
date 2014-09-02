package uvr

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "time"
    _"log"
    _"fmt"
)

func writeBits(bits []int, out chan bool, t timeout) {
    for _, b := range bits {
        if b == 0 {
            out <- false
        } else {
            out <- true
        }
        time.Sleep(t.duration)
    }
}

func TestSendingBytes(t *testing.T) {
    in := make(chan bool)
    
    signal := NewSignal(in)
    syncBitsTimeout := NewTimeout(488.0, 0.5)
    syncDecoder := NewSyncDecoder(signal.out, syncBitsTimeout)
    bitsTimeout := NewTimeout(244.0, 0.3)
    byteDecoder := NewByteDecoder(syncDecoder.out, bitsTimeout)
    packetDecoder := NewPacketDecoder(byteDecoder.out, 2)
    
    go func() {
         writeBits([]int{0,0,1,1,1,0,0,1,1,1,1,1,1,1,1}, in, syncBitsTimeout) // sync
         writeBits([]int{0,1,1,1,1,1,1,1,1,1}, in, bitsTimeout) // 0xFF
         writeBits([]int{0,1,0,1,0,1,0,1,0,1}, in, bitsTimeout) // 0x0F
    }()
    
    bytes := <- packetDecoder.out
    assert.NotNil(t, bytes)
    assert.Equal(t, len(bytes), 2)
    assert.Equal(t, bytes[0], Byte(0xFF))
    assert.Equal(t, bytes[1], Byte(0x55))
}