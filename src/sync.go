package uvr

import (
    "fmt"
    "math/big"
    "log"
)

type syncPattern struct {
    i int
    count int
    value big.Word
    timeout timeout
    last *Bit
}

type syncDecoder struct {
    synced bool
    pattern syncPattern
    in chan Bit
    out chan Bit
}

func NewSyncDecoder(in chan Bit, timeout timeout) syncDecoder {
    d := syncDecoder{}
    
    d.pattern = syncPattern{
                    count: 8, 
                    value: big.Word(1), 
                    timeout: timeout,
                }
    d.in = in
    d.out = make(chan Bit)
    go d.start()
    
    return d
}

func (s *syncDecoder) reset() {
    s.pattern.last = nil
    s.pattern.i = 0 // reset
}

func (s *syncDecoder) start() {
    for {
        select {
        case bit := <- s.in:
            if s.synced == true {
                s.out <- bit
            } else {
                pattern := s.pattern
                if pattern.last != nil {
                    last := pattern.last.Timestamp
                    if pattern.timeout.IsFutureSince(last) {
                        fmt.Println("Skipping")
                        break
                    } else if pattern.timeout.IsPastSince(last) {
                        fmt.Print("[", int(bit.Raw), "]")
                        s.reset()
                        break
                    } else if pattern.timeout.PlausibleSince(last) == false {
                        log.Fatal("Sync bits timeout error")
                    }
                }
                
                if bit.Raw == s.pattern.value {
                    fmt.Print(int(bit.Raw))
                    s.pattern.i++
                    s.pattern.last = &bit
                    if s.pattern.i == s.pattern.count {
                        s.synced = true
                        fmt.Println(", synced")
                    }
                } else {
                    s.reset()
                }
            }
        }
    }
}
