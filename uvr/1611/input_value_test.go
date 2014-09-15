package uvr1611

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "gouvr/uvr"
)

func TestTemperatureFromValue(t *testing.T) {
    value := uvr.NewValue([]uvr.Byte{
        uvr.Byte(0x50), // 0101 0000
        uvr.Byte(0x20), // 0010 0000
    })
    
    assert.Equal(t, InputValueToString(value), "8.0 °C")
}