package uvr

import (
    "github.com/kidoman/embd"
    "time"
)

type Reader struct {
    pin embd.DigitalPin
    pulseWidth time.Duration
    decoder *Decoder
}

func NewUVR1611Reader(pin embd.DigitalPin, pulseWidth time.Duration) Reader {
    r := Reader{pin: pin, pulseWidth:pulseWidth} 
    r.decoder = &Decoder{signalPulseWidth: pulseWidth}
   
    return r
}

func (r Reader) NextPacket() (UVR1611Packet, error) {
    var err error
    var value int
    decoder := r.decoder
    for (decoder.NumberOfPacketBytes() < 63) {
        value, err = r.pin.Read()
        if err != nil {
           break
        }
        
        bit := NewBitFromInt(value)
        decoder.ProcessBit(bit)
        time.Sleep(r.pulseWidth)
    }
    
    bytes := decoder.GetPacketBytes()
    return NewUVR1611Packet(bytes)
}