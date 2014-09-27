package uvr

import(
    "fmt"
    "math/big"
    "time"
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
    timeout Timeout
    last *Bit
}

type byteDecoder struct {
    BitConsumer
    SyncObserver
    consumer ByteConsumer
    
    encoding byteEncoding
    bits []Bit
}

func NewByteDecoder(consumer ByteConsumer, t Timeout) *byteDecoder {
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

func (d *byteDecoder) Reset() {
    d.consumer.Reset()
    d.reset()
}

func (d *byteDecoder) reset() {
    d.encoding.last = nil
    d.complete()    
}

func (d *byteDecoder) complete() {
    d.bits = make([]Bit, 0, cap(d.bits))
}

func (d *byteDecoder) SyncDone(t time.Time) {
    // d.encoding.last = &Bit{Raw:0, Timestamp:t}
}
    
func (d *byteDecoder) Consume(bit Bit) error {
    encoding := d.encoding
    if encoding.last != nil {
        delta := time.Duration(bit.Timestamp.UnixNano() - encoding.last.Timestamp.UnixNano()) 
        switch bit.CompareTimeoutToLast(encoding.timeout, *encoding.last) {
        case OrderedAscending:
            // fmt.Print("[", bit.Raw,"]")
            return nil // ignore
        case OrderedDescending:
            err := NewErrorf("[BYTE] Bit arrived too late %v", delta)
            // fmt.Println(err)
            return err
        case OrderedSame:
            
        }
    }
    
    // fmt.Print(bit.Raw)
    
    bits := append(d.bits, bit)
    d.encoding.last = &bit
    if len(bits) == cap(d.bits) {
        if encoding.start != nil {
            if bits[0].Raw != encoding.start.Raw {
                err := NewError("[BYTE] Start bit is wrong")
                fmt.Println(err)
                return err
            }
            bits = bits[1:]
        }
        
        if encoding.stop != nil {
            if bit.Raw != encoding.stop.Raw {
                err := NewError("[BYTE] Stop bit is wrong")
                fmt.Println(err)
                return err
            }
            bits = bits[:len(bits)]
        }
        
        b := ByteFromBits(bits)
        d.complete()
        d.consumer.Consume(b)
    } else {
        d.bits = bits
    }
    
    return nil
}
