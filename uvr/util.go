package uvr

import (
    "errors"
    "fmt"
)

func ByteFromBits(bits []Bit) Byte {
    var value uint8
    
    for index, bit := range bits {
        value |= uint8(bit.Raw) << uint(index)
    }
    
    return Byte(value)
}

func NewUVR1611Packet(bytes []Byte) (UVR1611Packet, error){
    var err error
    var packet = UVR1611Packet{}
    
    if len(bytes) == 64 {
        packet.deviceType = DeviceType(bytes[0])
        packet.deviceTypeInverted = DeviceType(bytes[1])
        packet.unkown = bytes[2]
        packet.timestamp = NewTimestamp(bytes[3:7])
        packet.sensor1 = NewValue(bytes[8:9])
        packet.sensor2 = NewValue(bytes[10:11])
        packet.sensor3 = NewValue(bytes[12:13])
        packet.sensor4 = NewValue(bytes[14:15])
        packet.sensor5 = NewValue(bytes[16:17])
        packet.sensor6 = NewValue(bytes[18:19])
        packet.sensor7 = NewValue(bytes[20:21])
        packet.sensor8 = NewValue(bytes[22:23])
        packet.sensor9 = NewValue(bytes[24:25])
        packet.sensor10 = NewValue(bytes[26:27])
        packet.sensor11 = NewValue(bytes[28:29])
        packet.sensor12 = NewValue(bytes[30:31])
        packet.sensor13 = NewValue(bytes[32:33])
        packet.sensor14 = NewValue(bytes[34:35])
        packet.sensor15 = NewValue(bytes[36:37])
        packet.sensor16 = NewValue(bytes[38:39])
        packet.outgoing = NewValue(bytes[40:41])
        packet.speedStepA1 = bytes[42]
        packet.speedStepA2 = bytes[43]
        packet.speedStepA6 = bytes[44]
        packet.speedStepA7 = bytes[45]
        packet.heatRegister = NewHeatMeterRegister(bytes[46])
        packet.heatMeter1 = NewHeatMeterValue(bytes[47:54])
        packet.heatMeter2 = NewHeatMeterValue(bytes[55:62])
        packet.checksum = bytes[63]
    } else {
        err = errors.New(fmt.Sprintf("Invalid number of uvr1611 packet bytes: %d", len(bytes)))
    }
    
    return packet, err
}