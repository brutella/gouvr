package uvr

import "math/big"

type WordConsumer interface {
    Consume(w big.Word) error
}

type BitConsumer interface {
    Consume(b Bit) error
}

type ByteConsumer interface {
    Consume(b Byte) error
}

type PacketConsumer interface {
    Consume(p Packet) error
}