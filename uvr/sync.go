package uvr

import (
    "fmt"
    "math/big"
    "time"
)

type SyncPattern struct {
    I int
    Count int
    Value big.Word
    Timeout Timeout
    Last *Bit
}

type syncDecoder struct {
    BitConsumer
    bitConsumer BitConsumer
    syncConsumer SyncConsumer
    synced bool
    pattern SyncPattern
}

func NewSyncDecoder(bitConsumer BitConsumer, syncConsumer SyncConsumer, t Timeout) *syncDecoder {
    d := &syncDecoder{bitConsumer: bitConsumer, syncConsumer: syncConsumer}
    
    d.pattern = SyncPattern{
                    Count: 8, 
                    Value: big.Word(1), 
                    Timeout: t,
                } 
                   
    return d
}

func (s *syncDecoder) resetBits() {
    s.pattern.Last = nil
    s.pattern.I = 0 // reset
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
        if pattern.Last != nil {
            delta := time.Duration(bit.Timestamp.UnixNano() - pattern.Last.Timestamp.UnixNano()) 
            switch bit.CompareTimeoutToLast(pattern.Timeout, *pattern.Last) {
            case OrderedAscending:
                fmt.Printf("[SYNC] Bit arrived too early (%v)\n", delta)
                return nil
            case OrderedDescending:
                s.Reset()
                err := NewErrorf("[SYNC] Bit arrived too late (%v)", delta)
                return err
            case OrderedSame:
            }
        }
        
        if bit.Raw == s.pattern.Value {
            s.pattern.I++
            s.pattern.Last = &bit
            if s.pattern.I == s.pattern.Count {
                if (s.syncConsumer != nil) {
                    s.syncConsumer.SyncDone(bit.Timestamp)
                }
                s.synced = true
                fmt.Println("[SYNC] Done")
            }
        } else {
            s.resetBits()
        }
    }
    return nil
}
