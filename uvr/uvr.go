package uvr

import (
    "github.com/kidoman/embd"
    "time"
    "errors"
)

type DeviceType Byte
type PollDelay time.Duration

const (
    DeviceTypeUVR1611 = DeviceType(0x80)
    DeviceTypeUVR61_3 = DeviceType(0x90)
)

const ClockUVR1611 = 488.0

type UVR struct {
    pin embd.DigitalPin
    clock float64 // Hz
    identifer uint8
}

func New(p embd.DigitalPin, t DeviceType) (UVR, error) {
    uvr := UVR{pin: p}
    
    err := p.SetDirection(embd.In)
    if err == nil {
        switch t {
        case DeviceTypeUVR1611:
            uvr.clock = ClockUVR1611
        default:
            err = errors.New("Could not create UVR other than Type1611")
        }
    }
        
    return uvr, err
}

func (u UVR) NextPacket() (UVR1611Packet, error) {
    pulseWidth := time.Duration(1/u.clock * float64(time.Second))
    r := NewUVR1611Reader(u.pin, pulseWidth)
    return r.NextPacket() 
}