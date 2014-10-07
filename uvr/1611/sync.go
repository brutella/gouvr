package uvr1611

import (
    _"fmt"
    "time"
    "math/big"
    "github.com/brutella/gouvr/uvr"
)

type syncDecoder struct {
    uvr.BitConsumer
    bitConsumer uvr.BitConsumer
    syncObserver uvr.SyncObserver
    synced bool
    pattern uvr.SyncPattern
}

func NewSyncDecoder(bitConsumer uvr.BitConsumer, syncObserver uvr.SyncObserver, t uvr.Timeout) *syncDecoder {
    d := &syncDecoder{bitConsumer: bitConsumer, syncObserver: syncObserver}
    
    d.pattern = uvr.SyncPattern{
                    Count: 32,
                    Timeout: t,
                } 
                   
    return d
}

func (s *syncDecoder) Reset() {
    s.bitConsumer.Reset()
    s.resetBits()
}

func (s *syncDecoder) resetBits() {
    s.pattern.Last = nil
    s.pattern.I = 0 // reset
    s.synced = false
}

func (s *syncDecoder) Consume(bit uvr.Bit) error {
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
            case uvr.OrderedAscending:
                // fmt.Printf("[SYNC] Bit arrived too early (%v)\n", delta)
                return nil
            case uvr.OrderedDescending:
                s.Reset()
                err := uvr.NewErrorf("[SYNC] Bit arrived too late (%v)", delta)
                return err
            case uvr.OrderedSame:
            }
        }
        
        if (pattern.Last == nil && bit.Raw == big.Word(0)) || (pattern.Last != nil && pattern.Last.Raw != bit.Raw) {
            s.pattern.I++
            s.pattern.Last = &bit
            if s.pattern.I == s.pattern.Count {
                if (s.syncObserver != nil) {
                    s.syncObserver.SyncDone(bit.Timestamp)
                }
                s.synced = true
                // fmt.Println("[SYNC] Done")
            }
        } else {
            s.Reset()
        }
    }
    return nil
}
