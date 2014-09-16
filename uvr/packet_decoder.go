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