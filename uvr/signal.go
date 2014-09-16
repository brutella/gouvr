package uvr

import(
    "math/big"
)

// Creates bit (w/ timestamp) stream out of a "raw" bit stream
type signal struct {
    WordConsumer
    consumer BitConsumer
}

func NewSignal(consumer BitConsumer) *signal {
    d := &signal{consumer: consumer}
    return d
}

func (s *signal) Consume(w big.Word) error {
    bit := NewBitFromWord(w)
    s.consumer.Consume(bit)
    return nil
}