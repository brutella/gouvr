package uvr

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestUVR1611PacketFromBytes(t *testing.T) {
    /*
            deviceType DeviceType           // 1
            deviceTypeInverted DeviceType   // 2
            unkown Byte                     // 3
            timestamp Timestamp             // 4, 5, 6, 7, 8
            sensor1 Value        // 9, 10
            sensor2 Value        // 11, 12
            sensor3 Value        // 13, 14
            sensor4 Value        // 15, 16
            sensor5 Value        // 17, 18
            sensor6 Value        // 19, 20
            sensor7 Value        // 21, 22
            sensor8 Value        // 23, 24
            sensor9 Value        // 25, 26
            sensor10 Value       // 27, 28
            sensor11 Value       // 29, 30
            sensor12 Value       // 31, 32
            sensor13 Value       // 33, 34
            sensor14 Value       // 35, 36
            sensor15 Value       // 37, 38
            sensor16 Value       // 39, 40
            outgoing Value        // 41, 42
            speedStepA1 Byte     // 43
            speedStepA2 Byte     // 44
            speedStepA6 Byte     // 45
            speedStepA7 Byte     // 46
            heatRegister HeatMeterRegister // 47
            heatMeter1 HeatMeterValue // 48, 49, 50, 51 | 52, 53 | 54, 55
            heatMeter2 HeatMeterValue // 56, 57, 58, 59 | 60, 61 | 62, 63
            checksum Byte       // 64
    */
    bytes := []Byte{ 
        Byte(0), 
    }
    packet, err := NewUVR1611Packet(bytes)
    assert.Nil(t, err)
    assert.Equal(t, packet.deviceType, DeviceTypeUVR1611)
}
