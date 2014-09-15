package uvr

import (
    "fmt"
    "time"
    "math/big"
)

type uvr1611SyncDecoder struct {
    BitConsumer
    bitConsumer BitConsumer
    syncConsumer SyncConsumer
    synced bool
    pattern syncPattern
}

func NewUVR1611SyncDecoder(bitConsumer BitConsumer, syncConsumer SyncConsumer, t timeout) *uvr1611SyncDecoder {
    d := &uvr1611SyncDecoder{bitConsumer: bitConsumer, syncConsumer: syncConsumer}
    
    d.pattern = syncPattern{
                    count: 32,
                    timeout: t,
                } 
                   
    return d
}

func (s *uvr1611SyncDecoder) Reset() {
    s.bitConsumer.Reset()
    s.resetBits()
}

func (s *uvr1611SyncDecoder) resetBits() {
    s.pattern.last = nil
    s.pattern.i = 0 // reset
    s.synced = false
}

func (s *uvr1611SyncDecoder) Consume(bit Bit) error {
    if s.synced == true {
        // bitConsumer returns error when bit order is wrong
        // e.g. wrong start/stop bit
        err := s.bitConsumer.Consume(bit)
        if err != nil {
            s.Reset()
        }
    } else {
        pattern := s.pattern
        if pattern.last != nil {
            delta := time.Duration(bit.Timestamp.UnixNano() - pattern.last.Timestamp.UnixNano()) 
            switch bit.CompareTimeoutToLast(pattern.timeout, *pattern.last) {
            case OrderedAscending:
                fmt.Printf("Bit arrived too early %v", delta)
                return nil
            case OrderedDescending:
                s.Reset()
                return NewErrorf("Bit arrived too late %v", delta)
            case OrderedSame:
            }
        }
        
        if (pattern.last == nil && bit.Raw == big.Word(0)) || (pattern.last != nil && pattern.last.Raw != bit.Raw) {
            s.pattern.i++
            s.pattern.last = &bit
            if s.pattern.i == s.pattern.count {
                if (s.syncConsumer != nil) {
                    s.syncConsumer.SyncDone(bit.Timestamp)
                }
                s.synced = true
                fmt.Println("Synced")
            }
        } else {
            s.Reset()
        }
    }
    return nil
}
