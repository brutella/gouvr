package uvr1611

// Consumers UVR1611 packets
type PacketConsumer interface {
    Consume(p Packet) error
}