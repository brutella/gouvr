package uvr

import(
    "fmt"
    "math/big"
    "log"
)
type byteEncoding struct {
    start *Bit
    stop *Bit
    timeout timeout
    last *Bit
}

type byteDecoder struct {
    in chan Bit
    out chan Byte
    
    encoding byteEncoding
    bits []Bit
}

func NewByteDecoder(in chan Bit, t timeout) byteDecoder {
    d := byteDecoder{}
    d.in = in
    d.out = make(chan Byte)
    
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
    
    go d.start()
    
    return d
}
func (d *byteDecoder) resetBits() {
    d.bits = make([]Bit, 0, cap(d.bits))
    d.encoding.last = nil
}
    
func (d *byteDecoder) start() {
    for {
        select {
        case bit := <- d.in:
            encoding := d.encoding
            if encoding.last != nil {
                last := encoding.last.Timestamp
                if encoding.timeout.IsFutureSince(last) {
                    fmt.Println("Skipping bit")
                    break
                } else if encoding.timeout.IsPastSince(last) {
                    log.Fatal("Bit arrived too late")
                    break
                } else if encoding.timeout.PlausibleSince(last) == false {
                    log.Fatal("Timeout error")
                }
            }
            
            fmt.Print(int(bit.Raw))
            bits := append(d.bits, bit)
            if len(bits) == cap(d.bits) {
                if encoding.start != nil {
                    if bits[0].Raw != encoding.start.Raw {
                        fmt.Println(", start bit is wrong")
                        break
                    }
                    bits = bits[1:]
                }
                
                if encoding.stop != nil {
                    if bit.Raw != encoding.stop.Raw {
                        fmt.Println(", stop bit is wrong")
                        break
                    }
                    bits = bits[:len(bits)]
                }
                
                // FIX index
                b := ByteFromBits(bits)
                fmt.Println(", new byte", b)
                d.out <- b
                d.resetBits()
            } else {
                d.bits = bits
                d.encoding.last = &bit
            }
        }
    }
}
