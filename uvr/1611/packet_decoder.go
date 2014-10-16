package uvr1611

import (
    "github.com/brutella/gouvr/uvr"
    "fmt"
)

type packetDecoder struct {
    uvr.ByteConsumer
    consumer PacketConsumer
    
    bytes []uvr.Byte
}

func NewPacketDecoder(consumer PacketConsumer) *packetDecoder {
    d := &packetDecoder{consumer: consumer}
    d.bytes = make([]uvr.Byte, 0, PacketByteCount)
    
    return d
}

func (d *packetDecoder) Reset() {
    d.bytes = make([]uvr.Byte, 0, cap(d.bytes))
}

func (d *packetDecoder) Consume(b uvr.Byte) error {
    bytes := append(d.bytes, b)
    if len(bytes) == cap(d.bytes) {
        d.Reset()
        packet, err := PacketFromBytes(bytes)
        if err != nil {
            fmt.Println("[PACKET] Could not parse packet bytes.", err)
            return err
        } else {
            d.consumer.Consume(packet)
        }
    } else {
        d.bytes = bytes
    }
    
    return nil
}