package uvr

import(
    "math/big"
)
type signal struct {
    in chan bool
    out chan Bit
}

func NewSignal(in chan bool) signal {
    d := signal{}
    d.in = in
    d.out = make(chan Bit)
    go d.start()
    
    return d
}

func (s *signal) start() {
    for {
        select {
        case b := <- s.in:
            word := big.Word(0)
            if b == true {word = big.Word(1)}
            bit := NewBitFromWord(word)
            s.out <- bit
        }
    }
}