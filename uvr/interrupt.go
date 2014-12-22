package uvr

import (
	_ "fmt"
	_ "time"
)

type edgeInterrupt struct {
	BitConsumer
	consumer BitConsumer
	last     *Bit
}

// NewEdgeInterrupt returns a bit consumer which forwards only edge changes
// from a bit stream to a bit consumer similar to an interrupt
func NewEdgeInterrupt(consumer BitConsumer) *edgeInterrupt {
	return &edgeInterrupt{consumer: consumer}
}

func (i *edgeInterrupt) Reset() {
}

func (i *edgeInterrupt) Consume(bit Bit) error {
	if i.last != nil && (i.last).Raw != bit.Raw {
		i.consumer.Consume(bit)
	}

	i.last = &bit

	return nil
}
