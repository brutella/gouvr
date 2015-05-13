package uvr

import (
	"math/big"
	"testing"
)

type bitConsumer struct {
	bits []Bit
}

func NewBitConsumer() *bitConsumer {
	return &bitConsumer{
		bits: []Bit{},
	}
}

func (c *bitConsumer) Consume(b Bit) error {
	c.bits = append(c.bits, b)
	return nil
}

func (c *bitConsumer) Reset() {
	c.bits = []Bit{}
}

func TestSignal(t *testing.T) {
	c := NewBitConsumer()
	s := NewSignal(c)
	s.Consume(big.Word(0))
	s.Consume(big.Word(1))

	if is, want := len(c.bits), 2; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	b0 := c.bits[0]
	b1 := c.bits[1]

	if is, want := b0.Raw, big.Word(0); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := b1.Raw, big.Word(1); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
