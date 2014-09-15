package uvr

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestTemperatureFromValue(t *testing.T) {
    value := NewValue([]Byte{
        Byte(0x50), // 0101 0000
        Byte(0x20), // 0010 0000
    })
    
    assert.Equal(t, InputValueToString(value), "8.0 °C")
}