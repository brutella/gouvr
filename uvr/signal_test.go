package uvr

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "math/big"
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
    
    assert.Equal(t, len(c.bits), 2)
    b0 := c.bits[0]
    b1 := c.bits[1]
    assert.Equal(t, b0.Raw, big.Word(0))
    assert.Equal(t, b1.Raw, big.Word(1))
}