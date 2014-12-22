package uvr1611

// PacketConsumer is an interface to consume UVR1611 packets
type PacketConsumer interface {
    Consume(p Packet) error
}