package uvr

import (
    "fmt"
    "math/big"
    "time"
)

type syncPattern struct {
    i int
    count int
    value big.Word
    timeout timeout
    last *Bit
}

type syncDecoder struct {
    BitConsumer
    bitConsumer BitConsumer
    syncConsumer SyncConsumer
    synced bool
    pattern syncPattern
}

func NewSyncDecoder(bitConsumer BitConsumer, syncConsumer SyncConsumer, t timeout) *syncDecoder {
    d := &syncDecoder{bitConsumer: bitConsumer, syncConsumer: syncConsumer}
    
    d.pattern = syncPattern{
                    count: 8, 
                    value: big.Word(1), 
                    timeout: t,
                } 
                   
    return d
}

func (s *syncDecoder) resetBits() {
    s.pattern.last = nil
    s.pattern.i = 0 // reset
    s.synced = false
}
func (s *syncDecoder) Reset() {
    s.bitConsumer.Reset()
    s.resetBits()
}

func (s *syncDecoder) Consume(bit Bit) error {
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
            println(delta)
            switch bit.CompareTimeoutToLast(pattern.timeout, *pattern.last) {
            case OrderedAscending:
                fmt.Println(NewErrorf("Bit arrived too early %v", delta))
                return nil
            case OrderedDescending:
                fmt.Println(NewErrorf("Bit arrived too late %v", delta))
            case OrderedSame:
            }
        }
        
        if bit.Raw == s.pattern.value {
            s.pattern.i++
            s.pattern.last = &bit
            if s.pattern.i == s.pattern.count {
                if (s.syncConsumer != nil) {
                    s.syncConsumer.SyncDone(bit.Timestamp)
                }
                s.synced = true
                fmt.Println("*Synced*")
            }
        } else {
            s.resetBits()
        }
    }
    return nil
}
