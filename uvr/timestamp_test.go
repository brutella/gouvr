package uvr

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestTimestampToString(t *testing.T) {
    bytes := []Byte{
        Byte(45), // min
        Byte(23), // h
        Byte(2), // d
        Byte(12), // M
        Byte(14), // y
    }
    timestamp := NewTimestamp(bytes)
    assert.Equal(t, timestamp.ToString(), "2014-12-02 23:45:00")
}