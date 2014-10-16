package uvr1611

// Consumers of UVR1611 packets
type PacketConsumer interface {
    Consume(p Packet) error
}