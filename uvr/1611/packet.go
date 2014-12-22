package uvr1611

import (
    "github.com/brutella/gouvr/uvr"
    "fmt"
    "errors"
)

const PacketByteCount = 64

// Packet represents an UVR1611 packet.
type Packet struct {                   // Bytes
    DeviceType uvr.DeviceType          // 1
    DeviceTypeInverted uvr.DeviceType  // 2
    Unkown uvr.Byte                    // 3
    Timestamp uvr.Timestamp            // 4, 5, 6, 7, 8
    Input1 uvr.Value                   // 9, 10 
    Input2 uvr.Value                   // 11, 12
    Input3 uvr.Value                   // 13, 14
    Input4 uvr.Value                   // 15, 16
    Input5 uvr.Value                   // 17, 18
    Input6 uvr.Value                   // 19, 20
    Input7 uvr.Value                   // 21, 22
    Input8 uvr.Value                   // 23, 24
    Input9 uvr.Value                   // 25, 26
    Input10 uvr.Value                  // 27, 28
    Input11 uvr.Value                  // 29, 30
    Input12 uvr.Value                  // 31, 32
    Input13 uvr.Value                  // 33, 34
    Input14 uvr.Value                  // 35, 36
    Input15 uvr.Value                  // 37, 38
    Input16 uvr.Value                  // 39, 40
    Outgoing uvr.Value                 // 41, 42
    SpeedStepA1 uvr.Byte               // 43
    SpeedStepA2 uvr.Byte               // 44
    SpeedStepA6 uvr.Byte               // 45
    SpeedStepA7 uvr.Byte               // 46
    HeatRegister uvr.HeatMeterRegister // 47
    HeatMeter1 uvr.HeatMeterValue      // 48, 49, 50, 51 | 52, 53 | 54, 55
    HeatMeter2 uvr.HeatMeterValue      // 56, 57, 58, 59 | 60, 61 | 62, 63
    Checksum uvr.Byte                  // 64
}

// Logs the packet to console.
func (p *Packet) Log() {
    fmt.Println("Device:", uvr.DeviceTypeToString(p.DeviceType))
    fmt.Println("Time:", p.Timestamp.ToString())
    fmt.Println("input 1:", InputValueToString(p.Input1))
    fmt.Println("input 2:", InputValueToString(p.Input2))
    fmt.Println("input 3:", InputValueToString(p.Input3))
    fmt.Println("input 4:", InputValueToString(p.Input4))
    fmt.Println("input 5:", InputValueToString(p.Input5))
    fmt.Println("input 6:", InputValueToString(p.Input6))
    fmt.Println("input 7:", InputValueToString(p.Input7))
    fmt.Println("input 8:", InputValueToString(p.Input8))
    fmt.Println("input 9:", InputValueToString(p.Input9))
    fmt.Println("input 10:", InputValueToString(p.Input10))
    fmt.Println("input 11:", InputValueToString(p.Input11))
    fmt.Println("input 12:", InputValueToString(p.Input12))
    fmt.Println("input 13:", InputValueToString(p.Input13))
    fmt.Println("input 14:", InputValueToString(p.Input14))
    fmt.Println("input 15:", InputValueToString(p.Input15))
    fmt.Println("input 16:", InputValueToString(p.Input16))
    fmt.Println("Out:")
    fmt.Println(OutletsStringFromValue(p.Outgoing))
    fmt.Println("A1:", uvr.NewSpeedStep(p.SpeedStepA1).ToString())
    fmt.Println("A2:", uvr.NewSpeedStep(p.SpeedStepA2).ToString())
    fmt.Println("A6:", uvr.NewSpeedStep(p.SpeedStepA6).ToString())
    fmt.Println("A7:", uvr.NewSpeedStep(p.SpeedStepA7).ToString())
        
    h1, h2 := AreHeatMetersEnabled(p.HeatRegister)
    if h1 == true {
        fmt.Println("Heat Meter 1:", p.HeatMeter1.ToString())
    }
    
    if h2 == true {
        fmt.Println("Heat Meter 2:", p.HeatMeter2.ToString())
    }
}

// Returns the sum of all bytes modulo 256.
func ChecksumFromBytes(bytes []uvr.Byte) uvr.Byte {
    var checksum uvr.Byte
    
    for _, b := range bytes {
        checksum += b
    }
    
    return checksum
}

// Returns a uvr1611 packet from a list of bytes or an error.
func PacketFromBytes(bytes []uvr.Byte) (Packet, error){
    var err error
    var packet = Packet{}
    
    if len(bytes) == PacketByteCount {
        packet.DeviceType = uvr.DeviceType(bytes[0])
        packet.DeviceTypeInverted = uvr.DeviceType(bytes[1])
        packet.Unkown = bytes[2]
        packet.Timestamp = uvr.NewTimestamp(bytes[3:8])
        packet.Input1 = uvr.NewValue(bytes[8:10])
        packet.Input2 = uvr.NewValue(bytes[10:12])
        packet.Input3 = uvr.NewValue(bytes[12:14])
        packet.Input4 = uvr.NewValue(bytes[14:16])
        packet.Input5 = uvr.NewValue(bytes[16:18])
        packet.Input6 = uvr.NewValue(bytes[18:20])
        packet.Input7 = uvr.NewValue(bytes[20:22])
        packet.Input8 = uvr.NewValue(bytes[22:24])
        packet.Input9 = uvr.NewValue(bytes[24:26])
        packet.Input10 = uvr.NewValue(bytes[26:28])
        packet.Input11 = uvr.NewValue(bytes[28:30])
        packet.Input12 = uvr.NewValue(bytes[30:32])
        packet.Input13 = uvr.NewValue(bytes[32:34])
        packet.Input14 = uvr.NewValue(bytes[34:36])
        packet.Input15 = uvr.NewValue(bytes[36:38])
        packet.Input16 = uvr.NewValue(bytes[38:40])
        packet.Outgoing = uvr.NewValue(bytes[40:42])
        packet.SpeedStepA1 = bytes[42]
        packet.SpeedStepA2 = bytes[43]
        packet.SpeedStepA6 = bytes[44]
        packet.SpeedStepA7 = bytes[45]
        packet.HeatRegister = uvr.NewHeatMeterRegister(bytes[46])
        packet.HeatMeter1 = uvr.NewHeatMeterValue(bytes[47:55])
        packet.HeatMeter2 = uvr.NewHeatMeterValue(bytes[55:63])
        packet.Checksum = bytes[63]
        
        checksum := ChecksumFromBytes(bytes[0:63])
        if packet.Checksum != checksum {
            err = errors.New(fmt.Sprintf("Invalid checksum %d should be %d.", checksum, packet.Checksum))
        }
    } else {
        err = errors.New(fmt.Sprintf("Invalid number of uvr1611 packet bytes: %d.", len(bytes)))
    }
    
    return packet, err
}