package uvr1611

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/brutella/gouvr/uvr"
)

func TestPacket(t *testing.T) {
    bytes := []uvr.Byte{
        uvr.Byte(uvr.DeviceTypeUVR1611),
        uvr.Byte(0x00),
        uvr.Byte(0x00),
        uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00),
        uvr.Byte(0xFA), uvr.Byte(0x20),
        uvr.Byte(0xAF), uvr.Byte(0x20),
        uvr.Byte(0x11), uvr.Byte(0x20),
        uvr.Byte(0x22), uvr.Byte(0x20),
        uvr.Byte(0x33), uvr.Byte(0x20),
        uvr.Byte(0x44), uvr.Byte(0x20),
        uvr.Byte(0x55), uvr.Byte(0x20),
        uvr.Byte(0x00), uvr.Byte(0x00),
        uvr.Byte(0x00), uvr.Byte(0x00),
        uvr.Byte(0x00), uvr.Byte(0x00),
        uvr.Byte(0x00), uvr.Byte(0x00),
        uvr.Byte(0x00), uvr.Byte(0x00),
        uvr.Byte(0x00), uvr.Byte(0x00),
        uvr.Byte(0x00), uvr.Byte(0x00),
        uvr.Byte(0x00), uvr.Byte(0x00),
        uvr.Byte(0x00), uvr.Byte(0x00),
        uvr.Byte(0x55), uvr.Byte(0x55), // 0101 0101 0101 0101
        uvr.Byte(0x00),
        uvr.Byte(0x00),
        uvr.Byte(0x00),
        uvr.Byte(0x00),
        uvr.Byte(0x00),
        uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00),
        uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00), uvr.Byte(0x00),
    }
    bytes = append(bytes, ChecksumFromBytes(bytes))
    
    packet, err := PacketFromBytes(bytes)
    assert.Nil(t, err)
    assert.NotNil(t, packet)
    
    assert.Equal(t, packet.DeviceType, uvr.DeviceTypeUVR1611)
    assert.Equal(t, uvr.UInt16FromValue(packet.Input1), uint16(0x20FA))
    assert.Equal(t, uvr.UInt16FromValue(packet.Input2), uint16(0x20AF))
    assert.Equal(t, uvr.UInt16FromValue(packet.Input3), uint16(0x2011))
    assert.Equal(t, uvr.UInt16FromValue(packet.Input4), uint16(0x2022))
    assert.Equal(t, uvr.UInt16FromValue(packet.Input5), uint16(0x2033))
    assert.Equal(t, uvr.UInt16FromValue(packet.Input6), uint16(0x2044))
    assert.Equal(t, uvr.UInt16FromValue(packet.Input7), uint16(0x2055))
}