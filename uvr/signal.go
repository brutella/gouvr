package uvr

import(
    "math/big"
)

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