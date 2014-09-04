package uvr

import(
    "fmt"
    "math/big"
)

type Byte uint8

func ByteFromBits(bits []Bit) Byte {
    var value uint8
    for index, bit := range bits {
        value |= uint8(bit.Raw) << uint(index)
    }
    
    return Byte(value)
}

type byteEncoding struct {
    start *Bit
    stop *Bit
    timeout timeout
    last *Bit
}

type byteDecoder struct {
    BitConsumer
    consumer ByteConsumer
    
    encoding byteEncoding
    bits []Bit
}

func NewByteDecoder(consumer ByteConsumer, t timeout) *byteDecoder {
    d := &byteDecoder{consumer:consumer}
    
    encoding := byteEncoding{}
    encoding.start = new(Bit)
    encoding.start.Raw = big.Word(0)
    encoding.stop = new(Bit)
    encoding.stop.Raw = big.Word(1)
    encoding.timeout = t
    d.encoding = encoding
    
    bits := 8
    if encoding.start != nil { bits++ }
    if encoding.stop != nil { bits++ }
    d.bits = make([]Bit, 0, bits)
    
    return d
}
func (d *byteDecoder) resetBits() {
    d.bits = make([]Bit, 0, cap(d.bits))
    d.encoding.last = nil
}
    
func (d *byteDecoder) Consume(bit Bit) error {
    encoding := d.encoding
    if encoding.last != nil {
        switch bit.CompareTimeoutToLast(encoding.timeout, *encoding.last) {
        case OrderedAscending:
            fmt.Println("Skipping")
            return nil
        case OrderedDescending:
            err := NewErrorf("Bit arrival at %d is too late for timeout %d (+/- %f)", bit.Timestamp, encoding.timeout.duration, encoding.timeout.deviation)
            d.resetBits()
            return err
        case OrderedSame:
            // ok
        }
    }
    
    bits := append(d.bits, bit)
    if len(bits) == cap(d.bits) {
        if encoding.start != nil {
            if bits[0].Raw != encoding.start.Raw {
                d.resetBits()
                return NewErrorf("Start bit is wrong")
            }
            bits = bits[1:]
        }
        
        if encoding.stop != nil {
            if bit.Raw != encoding.stop.Raw {
                d.resetBits()
                return NewErrorf("Stop bit is wrong")
            }
            bits = bits[:len(bits)]
        }
        
        // FIX index
        b := ByteFromBits(bits)
        fmt.Printf("[%b]\n", b)
        d.resetBits()
        d.consumer.Consume(b)
    } else {
        d.bits = bits
        d.encoding.last = &bit
    }
    
    return nil
}
