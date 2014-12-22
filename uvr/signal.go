package uvr

import (
	"math/big"
)

type signal struct {
	WordConsumer
	consumer BitConsumer
}

// NewSignal returns a word consumer which creates a Bit stream out of a Word stream
func NewSignal(consumer BitConsumer) *signal {
	d := &signal{consumer: consumer}
	return d
}

func (s *signal) Consume(w big.Word) error {
	bit := NewBitFromWord(w)
	s.consumer.Consume(bit)
	return nil
}
