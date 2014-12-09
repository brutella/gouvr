package uvr1611

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/brutella/gouvr/uvr"
)

func TestTemperatureInputValue(t *testing.T) {
    value := uvr.NewValue([]uvr.Byte{
        uvr.Byte(0x50), // 0101 0000
        uvr.Byte(0x20), // 0010 0000
    })
    
    input_type, f := DecodeInputValue(value)
    assert.Equal(t, input_type, InputTypeTemperature)
    assert.Equal(t, f, 8.0)
}

func TestRoomTemperatureInputValue(t *testing.T) {
    value := uvr.NewValue([]uvr.Byte{
        uvr.Byte(0x50), // 0101 0000
        uvr.Byte(0x74), // 0111 0100
    })
    
    input_type, f := DecodeInputValue(value)
    assert.Equal(t, input_type, InputTypeRoomTemperature)
    assert.Equal(t, f, 8.0)
    
    mode := RoomTemperatureModeFromValue(value)
    assert.Equal(t, mode, RoomTemperatureModeLowering)
}

func TestDigitalInputValue(t *testing.T) {
    value := uvr.NewValue([]uvr.Byte{
        uvr.Byte(0x50), // 0101 0000
        uvr.Byte(0x90), // 1001 0000
    })
    
    assert.True(t, IsDigitalInputValueOn(value))
}

func TestUnusedInputValue(t *testing.T) {
    value := uvr.NewValue([]uvr.Byte{
        uvr.Byte(0x10), // 0001 0000
        uvr.Byte(0x00), // 0000 0000
    })
    
    assert.True(t, IsUnusedInputValue(value))
}