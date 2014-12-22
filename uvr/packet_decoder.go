package uvr

// Packet encoding indicates the number of bytes required for a packet
type packetEncoding struct {
    count int
}

type packetDecoder struct {
    ByteConsumer
    
    consumer PacketConsumer
    
    encoding packetEncoding
    bytes []Byte
}

// NewPacketDecoder returns byte consumer which creates a packet from a list of bytes.
// The PacketConsumer's Consume method is called for every new packet.
// The byte_count argument specifies the number of bytes in a packet.
func NewPacketDecoder(consumer PacketConsumer, byte_count int) *packetDecoder {
    d := &packetDecoder{consumer: consumer}
    d.encoding = packetEncoding{count: byte_count}
    d.bytes = make([]Byte, 0, byte_count)
    
    return d
}

func (d *packetDecoder) Reset() {
    d.resetBytes()
}
    
func (d *packetDecoder) Consume(b Byte) error {
    bytes := append(d.bytes, b)
    if len(bytes) == cap(d.bytes) {
        d.resetBytes()
        d.consumer.Consume(Packet(bytes))
    } else {
        d.bytes = bytes
    }
    
    return nil
}

func (d *packetDecoder) resetBytes() {
    d.bytes = make([]Byte, 0, cap(d.bytes))
}