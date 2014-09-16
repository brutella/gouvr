package uvr

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "time"
    "math/big"
)

func Microseconds(t time.Time) int64 {
    return int64((t.UnixNano() % 1e9) / 1e3)
}

func Milliseconds(t time.Time) int64 {
    return int64((t.UnixNano() % 1e9) / 1e6)
}

func TestLogStringFromBit(t *testing.T) {
    bit := NewBitFromWord(1)
    
    str := LogString(bit)
    assert.NotNil(t, str)
    
    restored_bit, err := BitFromLogString(str)
    assert.Nil(t, err)
    assert.Equal(t, bit.Raw, restored_bit.Raw)
    
    // Only test milliseconds
    assert.Equal(t, Milliseconds(bit.Timestamp), Milliseconds(restored_bit.Timestamp))
}

func TestBitsWithinTimeout(t *testing.T) {
    bit1 := Bit{Raw:big.Word(1), Timestamp:time.Unix(0, 1001)}
    bit2 := Bit{Raw:big.Word(0), Timestamp:time.Unix(0, 1002)}
    
    timeout := Timeout{duration:time.Duration(1), deviation:0}
    // time difference is 1002 - 1001 = 1
    // valid timeout = +/-1
    assert.Equal(t, bit2.CompareTimeoutToLast(timeout, bit1), OrderedSame)
}

func TestBitsWithinTimeout2(t *testing.T) {
    bit1 := Bit{Raw:big.Word(1), Timestamp:time.Unix(0, 1001)}
    bit2 := Bit{Raw:big.Word(0), Timestamp:time.Unix(0, 1003)}
    
    timeout := Timeout{duration:time.Duration(1), deviation:1}
    // time difference is 1003 - 1001 = 2
    // valid timeout = +/-2
    assert.Equal(t, bit2.CompareTimeoutToLast(timeout, bit1), OrderedSame)
}

func TestBitsOutOfTimeout(t *testing.T) {
    bit1 := Bit{Raw:big.Word(1), Timestamp:time.Unix(0, 1001)}
    bit2 := Bit{Raw:big.Word(0), Timestamp:time.Unix(0, 1003)}
    
    timeout := Timeout{duration:time.Duration(1), deviation:0}
    // time difference is 1003 - 1001 = 2
    // valid timeout = +/-1
    assert.Equal(t, bit2.CompareTimeoutToLast(timeout, bit1), OrderedDescending)
}