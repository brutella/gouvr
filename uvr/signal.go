package uvr

import(
    "fmt"
    "errors"
    "time"
)

type Signal interface {
    IsComplete() bool
    Clear()
    GetBytes() ([]Byte, error)
    AddBit(bit Bit) error
}

type SyncSignal struct {
    bits []Bit
    lastBit *Bit
    pulseWidth time.Duration
}

func NewSyncSignal(capacity int, pulseWidth time.Duration) *SyncSignal {
    s := &SyncSignal{}
    s.pulseWidth = pulseWidth
    s.bits = make([]Bit, 0, capacity)
    
    return s
}

func (s *SyncSignal) Clear() {
    s.bits = make([]Bit, 0, cap(s.bits))
}

func (s *SyncSignal) IsComplete() bool {
    return len(s.bits) == cap(s.bits)
}

func (s *SyncSignal) GetBytes() ([]Byte, error) {
    if s.IsComplete() {
        lowByte := ByteFromBits(s.bits[0:8])
        highByte := ByteFromBits(s.bits[8:16])
        value := []Byte{lowByte, highByte}
        return value, nil
    }
    
    return nil, errors.New(fmt.Sprintf("Sync byte is of length %d must be of length %d", len(s.bits), cap(s.bits)))
}

func (s *SyncSignal) AddBit(bit Bit) error {
    if s.lastBit != nil {
        duration := time.Since(s.lastBit.Timestamp)
        r := NewRangeFromDuration(s.pulseWidth, 0.3)
        if DurationWithinRange(duration, r) == false {
            fmt.Println("Skipping bit")
            return errors.New("New bit within pulse width not allowed")
        }
    }
    
    var err error
    if len(s.bits) < cap(s.bits) {
        s.bits = append(s.bits, bit)
        s.lastBit = &bit
    } else {
        err = errors.New("Sync byte already full")
    }
    
    return err
}

type DataSignal struct {
    start Bit
    stop Bit
    
    lastBit *Bit
    pulseWidth time.Duration
    bits []Bit
}

func NewDataSignal(capacity int, pulseWidth time.Duration) *DataSignal {
    s := &DataSignal{}
    s.start = NewBitFromInt(0)
    s.stop = NewBitFromInt(1)
    
    s.pulseWidth = pulseWidth
    s.bits = make([]Bit, 0, capacity)
    
    return s
}

func (s *DataSignal) Clear() {
    s.bits = make([]Bit, 0, cap(s.bits))
}

func (d *DataSignal) IsComplete() bool {
    return len(d.bits) == cap(d.bits)
}

func (d *DataSignal) GetBytes() ([]Byte, error) {
    var value []Byte
    var err error
    if d.IsComplete() {
        if d.bits[0].Raw != d.start.Raw {
            err = errors.New("Start bit is not 0")
        } else if d.bits[9].Raw != d.stop.Raw {
            err = errors.New("Stop bit is not 1")
        } else {
            byteValue := ByteFromBits(d.bits[1:9])
            value = []Byte{byteValue}
        }
    } else {
        err = errors.New(fmt.Sprintf("Data byte is of length %d must be of length %d", len(d.bits), cap(d.bits)))
    }
    
    return value, err 
}

func (b *DataSignal) AddBit(bit Bit) error {
    if b.lastBit != nil {
        duration := time.Since(b.lastBit.Timestamp)
        r := NewRangeFromDuration(b.pulseWidth, 0.5)
        if DurationWithinRange(duration, r) == false {
            fmt.Println("Skipping bit")
            return errors.New("New bit within pulse width not allowed")
        }
    }
    
    var err error
    if len(b.bits) < cap(b.bits) {
        b.bits = append(b.bits, bit)
        b.lastBit = &bit
    } else {
        err = errors.New("Data byte already full")
    }
    
    return err
}
