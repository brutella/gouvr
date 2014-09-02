package uvr

import (
    "time"
    "math/big"
)

type Bit struct {
    Raw big.Word
    Timestamp time.Time
}

func NewBitFromWord(value big.Word) Bit {
    return Bit{Raw: value, Timestamp: time.Now()}
}

type Byte uint8

func ByteFromBits(bits []Bit) Byte {
    var value uint8
    for index, bit := range bits {
        value |= uint8(bit.Raw) << uint(index)
    }
    
    return Byte(value)
}
