package uvr

import (
    "fmt"
    "math/big"
    "errors"
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
    consumer BitConsumer
    synced bool
    pattern syncPattern
}

func NewSyncDecoder(consumer BitConsumer, t timeout) *syncDecoder {
    d := &syncDecoder{consumer: consumer}
    
    d.pattern = syncPattern{
                    count: 8, 
                    value: big.Word(1), 
                    timeout: t,
                } 
                   
    return d
}

func (s *syncDecoder) reset() {
    s.pattern.last = nil
    s.pattern.i = 0 // reset
    s.synced = false
}

func (s *syncDecoder) Consume(bit Bit) error {
    if s.synced == true {
        // consumer returns error when bit order is wrong
        // e.g. wrong start/stop bit
        err := s.consumer.Consume(bit)
        if err != nil {
            s.reset()
        }
    } else {
        pattern := s.pattern
        if pattern.last != nil {
            switch bit.CompareTimeoutToLast(pattern.timeout, *pattern.last) {
            case OrderedAscending:
                fmt.Println("Skipping")
                return nil
            case OrderedDescending:
                err := errors.New(fmt.Sprintf("Bit arrival at %d (delta %d)is too late for timeout %d (+/- %f)", bit.Timestamp.UnixNano(), bit.Timestamp.UnixNano() - pattern.last.Timestamp.UnixNano(), pattern.timeout.duration, pattern.timeout.deviation))
                fmt.Println(err)
                s.reset()
                return err
            case OrderedSame:
            }
        }
        
        if bit.Raw == s.pattern.value {
            s.pattern.i++
            s.pattern.last = &bit
            if s.pattern.i == s.pattern.count {
                s.synced = true
                fmt.Println("*Synced*")
            }
        } else {
            s.reset()
        }
    }
    return nil
}
