package uvr

import "fmt"

type Packet struct {
    byte1 Byte
    byte2 Byte
}

type packetEncoding struct {
    count int
}

type packetDecoder struct {
    in chan Byte
    out chan []Byte
    
    encoding packetEncoding
    bytes []Byte
}

func NewPacketDecoder(in chan Byte, byte_count int) packetDecoder {
    d := packetDecoder{}
    d.in = in
    d.out = make(chan []Byte)
    d.encoding = packetEncoding{count: byte_count}
    d.bytes = make([]Byte, 0, byte_count)
    
    go d.start()
    
    return d
}

func (d *packetDecoder) resetBytes() {
    d.bytes = make([]Byte, 0, cap(d.bytes))
}

func (d *packetDecoder) start() {
    for {
        select {
        case b := <- d.in:
            bytes := append(d.bytes, b)
            if len(bytes) == cap(d.bytes) {
                fmt.Println("packet ", bytes)
                d.out <- bytes
                d.resetBytes()
            } else {
                d.bytes = bytes
            }
        }
    }
}