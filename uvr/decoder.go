package uvr

import(
    "fmt"
    "time"
    _"errors"
)

type Decoder struct {
    packetBytes []Byte
    synced bool
    signal Signal
    signalPulseWidth time.Duration
}

func (d *Decoder) ProcessBit(b Bit) {
    if d.synced == false {
        if d.signal == nil {
            d.signal = NewSyncSignal(16, d.signalPulseWidth)
        }
        
        if (b.Raw != 1) {
            d.signal.Clear()
        } else {
            d.signal.AddBit(b)
        }
        
        if d.signal.IsComplete() {
            var bytes []Byte
            bytes, _ = d.signal.GetBytes()
            if bytes[0] == Byte(0xFF) && bytes[1] == Byte(0xFF) {
                d.synced = true
            } else {
                fmt.Println("Sync singal invalid", bytes)
            }
            // reset
            d.signal = nil
        }
    } else {
        if d.signal == nil {
            d.signal = NewDataSignal(10, 2*d.signalPulseWidth)
        }
        
        d.signal.AddBit(b)
        if d.signal.IsComplete() {
            var bytes []Byte
            bytes, _ = d.signal.GetBytes()
            if len(bytes) == 1 {
                d.packetBytes = append(d.packetBytes, bytes...)
            } else {
                fmt.Println("Data singal invalid", bytes)
            }
            // reset
            d.signal = nil
        }
    }
}

func (d *Decoder) GetPacketBytes() ([]Byte){
    return d.packetBytes
}

func (d *Decoder) NumberOfPacketBytes() int {
    return len(d.packetBytes)
}

func (d *Decoder) IsSynced() bool {
    return d.synced
}
