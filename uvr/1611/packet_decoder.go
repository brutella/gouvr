package uvr1611

import (
	"fmt"
	"github.com/brutella/gouvr/uvr"
)

type packetDecoder struct {
	uvr.ByteConsumer
	consumer PacketConsumer

	bytes []uvr.Byte
}

// NewPacketDecoder returns a new packet decoder, which implements the ByteConsumer interface.
// The decoder calls the Consume() method of the specified packet consumer when new packets were decoded
func NewPacketDecoder(consumer PacketConsumer) *packetDecoder {
	d := &packetDecoder{consumer: consumer}
	d.bytes = make([]uvr.Byte, 0, PacketByteCount)

	return d
}

// Resets the cached bytes.
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
