package uvr1611

// Consumers UVR1611 packes
type PacketConsumer interface {
    Consume(p Packet) error
}