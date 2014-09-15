package uvr

import (
    "math/big"
    "time"
)

type WordConsumer interface {
    Consume(w big.Word) error
}

type BitConsumer interface {
    Consume(b Bit) error
    Reset()
}

type SyncConsumer interface {
    SyncDone(t time.Time)
}

type ByteConsumer interface {
    Consume(b Byte) error
    Reset()
}

type PacketConsumer interface {
    Consume(p Packet) error
}