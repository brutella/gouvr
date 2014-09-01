package uvr

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "time"
)

func NewSyncedDecoder() Decoder {
    signalPulseWidth := time.Duration(time.Second/ClockUVR1611)
    d := Decoder{signalPulseWidth: signalPulseWidth}    
    d.synced = true
            
    return d
}

func TestDecoderSync(t *testing.T) {
    signalPulseWidth := time.Duration(time.Second/ClockUVR1611)
    d := Decoder{signalPulseWidth: signalPulseWidth}    
    for i := 0; i < 16; i++ {
        assert.False(t, d.IsSynced())
        time.Sleep(d.signalPulseWidth)
        d.ProcessBit(NewBitFromInt(1))
    }
    
    assert.True(t, d.IsSynced())
}

func TestDecoderPacket(t *testing.T) {
    d := NewSyncedDecoder()
    d.ProcessBit(NewBitFromInt(0))
    time.Sleep(d.signalPulseWidth)
    assert.Equal(t, d.NumberOfPacketBytes(), 0)
    d.ProcessBit(NewBitFromInt(1))
    time.Sleep(d.signalPulseWidth)
    assert.Equal(t, d.NumberOfPacketBytes(), 0)
    d.ProcessBit(NewBitFromInt(1))
    time.Sleep(d.signalPulseWidth)
    assert.Equal(t, d.NumberOfPacketBytes(), 0)
    d.ProcessBit(NewBitFromInt(1))
    time.Sleep(d.signalPulseWidth)
    assert.Equal(t, d.NumberOfPacketBytes(), 0)
    d.ProcessBit(NewBitFromInt(1))
    time.Sleep(d.signalPulseWidth)
    assert.Equal(t, d.NumberOfPacketBytes(), 0)
    d.ProcessBit(NewBitFromInt(0))
    time.Sleep(d.signalPulseWidth)
    assert.Equal(t, d.NumberOfPacketBytes(), 0)
    d.ProcessBit(NewBitFromInt(0))
    time.Sleep(d.signalPulseWidth)
    assert.Equal(t, d.NumberOfPacketBytes(), 0)
    d.ProcessBit(NewBitFromInt(0))
    time.Sleep(d.signalPulseWidth)
    assert.Equal(t, d.NumberOfPacketBytes(), 0)
    d.ProcessBit(NewBitFromInt(0))
    time.Sleep(d.signalPulseWidth)
    assert.Equal(t, d.NumberOfPacketBytes(), 0)
    d.ProcessBit(NewBitFromInt(1))
    time.Sleep(d.signalPulseWidth)
    assert.Equal(t, d.NumberOfPacketBytes(), 1)
    
    bytes := d.GetPacketBytes()
    assert.Equal(t, bytes[0], Byte(0x0F))
}