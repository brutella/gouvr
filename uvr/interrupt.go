package uvr

import (
    _"fmt"
    _"time"
)

type edgeInterrupt struct {
    BitConsumer
    consumer BitConsumer
    last *Bit
}

func NewEdgeInterrupt(consumer BitConsumer) *edgeInterrupt {
    return &edgeInterrupt{consumer: consumer}
}

func (i *edgeInterrupt) Reset() {}

func (i *edgeInterrupt) Consume(bit Bit) error {    
    if(i.last != nil && (i.last).Raw != bit.Raw) {
        i.consumer.Consume(bit)
    }
    
    i.last = &bit
    
    return nil
}