package uvr

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "time"
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
