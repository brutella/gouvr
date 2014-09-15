package uvr

import(
    "math/big"
    "time"
    "os"
    "strconv"
    _"fmt"
)

type UVR1611PacketCallback func(packet UVR1611Packet)
// UVR1611 Packet receiver
type uvr1611PacketReceiver struct {
    UVR1611PacketConsumer
    Packet UVR1611Packet
    
    callback UVR1611PacketCallback
}

func (p *uvr1611PacketReceiver) Consume(packet UVR1611Packet) error {
    p.Packet = packet
    
    if p.callback != nil {
        p.callback(p.Packet)
    }
    
    return nil
}

func NewUVR1611PacketReceiver() *uvr1611PacketReceiver {
    return &uvr1611PacketReceiver{}
}

func (p *uvr1611PacketReceiver) RegisterCallback(callback UVR1611PacketCallback) {
    p.callback = callback
}

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

func writeWords(words []big.Word, c WordConsumer, t timeout) {
    for _, w := range words {
        time.Sleep(t.duration)
        c.Consume(w)        
    }
}

func RandomTempFilePath() string {
    return os.TempDir() + "/" + strconv.FormatInt(time.Now().Unix(), 10) + ".log"
}