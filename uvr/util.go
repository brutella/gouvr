package uvr

import (
	_ "fmt"
	"math/big"
	"os"
	"strconv"
	"time"
)

// Packet receiver
type packetReceiver struct {
	PacketConsumer
	packet Packet
}

func (p *packetReceiver) Consume(packet Packet) error {
	p.packet = packet
	return nil
}

// Byte receiver
type byteReceiver struct {
	ByteConsumer
	bytes []Byte
}

func NewTestByteReceiver() *byteReceiver {
	return &byteReceiver{bytes: make([]Byte, 0, 2)}
}

func (receiver *byteReceiver) Consume(b Byte) error {
	receiver.bytes = append(receiver.bytes, b)
	return nil
}

func (receiver *byteReceiver) Reset() {
	receiver.bytes = make([]Byte, 0, cap(receiver.bytes))
}

// Bit receiver
type bitReceiver struct {
	BitConsumer
	bits []Bit
}

func NewTestBitReceiver() *bitReceiver {
	return &bitReceiver{bits: make([]Bit, 0, 2)}
}

func (receiver *bitReceiver) Consume(b Bit) error {
	receiver.bits = append(receiver.bits, b)
	return nil
}

func writeBits(bits []Bit, c BitConsumer) {
	for _, b := range bits {
		c.Consume(b)
	}
}

func writeWords(words []big.Word, c WordConsumer, t Timeout) {
	for _, w := range words {
		time.Sleep(t.duration)
		c.Consume(w)
	}
}

func RandomTempFilePath() string {
	return os.TempDir() + "/" + strconv.FormatInt(time.Now().Unix(), 10) + ".log"
}
