package uvr

import (
    "fmt"
    "errors"
)

const UVR1611PacketByteCount = 64

type UVR1611Packet struct {
    deviceType DeviceType           // 1
    deviceTypeInverted DeviceType   // 2
    unkown Byte                     // 3
    timestamp Timestamp             // 4, 5, 6, 7, 8
    input1 Value        // 9, 10  (Aussentemperatur)
    input2 Value        // 11, 12 (Fussbodenheizung Vorlauf)
    input3 Value        // 13, 14 (Buffer Warmwasser - Oben)
    input4 Value        // 15, 16 (Buffer Warmwasser - Mitte)
    input5 Value        // 17, 18 (Buffer Warmwasser - Unten)
    input6 Value        // 19, 20 (Raumtemperatur)
    input7 Value        // 21, 22 (Wärmetauscher Sekundär)
    input8 Value        // 23, 24
    input9 Value        // 25, 26
    input10 Value       // 27, 28
    input11 Value       // 29, 30
    input12 Value       // 31, 32
    input13 Value       // 33, 34
    input14 Value       // 35, 36
    input15 Value       // 37, 38
    input16 Value       // 39, 40
    outgoing Value        // 41, 42
    speedStepA1 Byte     // 43
    speedStepA2 Byte     // 44
    speedStepA6 Byte     // 45
    speedStepA7 Byte     // 46
    heatRegister HeatMeterRegister // 47
    heatMeter1 HeatMeterValue // 48, 49, 50, 51 | 52, 53 | 54, 55
    heatMeter2 HeatMeterValue // 56, 57, 58, 59 | 60, 61 | 62, 63
    checksum Byte       // 64
}

func (p *UVR1611Packet) Log() {
    fmt.Println("Device Type:", p.deviceType)
    fmt.Println("Inverted Device Type:", p.deviceTypeInverted)
    fmt.Println("Time:", p.timestamp.toString())
    fmt.Println("input 1:", InputValueToString(p.input1))
    fmt.Println("input 2:", InputValueToString(p.input2))
    fmt.Println("input 3:", InputValueToString(p.input3))
    fmt.Println("input 4:", InputValueToString(p.input4))
    fmt.Println("input 5:", InputValueToString(p.input5))
    fmt.Println("input 6:", InputValueToString(p.input6))
    fmt.Println("input 7:", InputValueToString(p.input7))
    fmt.Println("input 8:", InputValueToString(p.input8))
    fmt.Println("input 9:", InputValueToString(p.input9))
    fmt.Println("input 10:", InputValueToString(p.input10))
    fmt.Println("input 11:", InputValueToString(p.input11))
    fmt.Println("input 12:", InputValueToString(p.input12))
    fmt.Println("input 13:", InputValueToString(p.input13))
    fmt.Println("input 14:", InputValueToString(p.input14))
    fmt.Println("input 15:", InputValueToString(p.input15))
    fmt.Println("input 16:", InputValueToString(p.input16))
    fmt.Println("Out:", Int16FromValue(p.outgoing))
    fmt.Println("A1:", NewSpeedStep(p.speedStepA1).toString())
    fmt.Println("A2:", NewSpeedStep(p.speedStepA2).toString())
    fmt.Println("A6:", NewSpeedStep(p.speedStepA6).toString())
    fmt.Println("A7:", NewSpeedStep(p.speedStepA7).toString())
    fmt.Println("Heat Register:", p.heatRegister)
    fmt.Println("Heat Meter 1:", p.heatMeter1.toString())
    fmt.Println("Heat Meter 2:", p.heatMeter2.toString())
}

type UVR1611PacketConsumer interface {
    Consume(p UVR1611Packet) error
}

func UVR1611PacketFromBytes(bytes []Byte) (UVR1611Packet, error){
    var err error
    var packet = UVR1611Packet{}
    
    if len(bytes) == UVR1611PacketByteCount {
        packet.deviceType = DeviceType(bytes[0])
        packet.deviceTypeInverted = DeviceType(bytes[1])
        packet.unkown = bytes[2]
        packet.timestamp = NewTimestamp(bytes[3:8])
        packet.input1 = NewValue(bytes[8:10])
        packet.input2 = NewValue(bytes[10:12])
        packet.input3 = NewValue(bytes[12:14])
        packet.input4 = NewValue(bytes[14:16])
        packet.input5 = NewValue(bytes[16:18])
        packet.input6 = NewValue(bytes[18:20])
        packet.input7 = NewValue(bytes[20:22])
        packet.input8 = NewValue(bytes[22:24])
        packet.input9 = NewValue(bytes[24:26])
        packet.input10 = NewValue(bytes[26:28])
        packet.input11 = NewValue(bytes[28:30])
        packet.input12 = NewValue(bytes[30:32])
        packet.input13 = NewValue(bytes[32:34])
        packet.input14 = NewValue(bytes[34:36])
        packet.input15 = NewValue(bytes[36:38])
        packet.input16 = NewValue(bytes[38:40])
        packet.outgoing = NewValue(bytes[40:42])
        packet.speedStepA1 = bytes[42]
        packet.speedStepA2 = bytes[43]
        packet.speedStepA6 = bytes[44]
        packet.speedStepA7 = bytes[45]
        packet.heatRegister = NewHeatMeterRegister(bytes[46])
        packet.heatMeter1 = NewHeatMeterValue(bytes[47:55])
        packet.heatMeter2 = NewHeatMeterValue(bytes[55:63])
        packet.checksum = bytes[63]
    } else {
        err = errors.New(fmt.Sprintf("Invalid number of uvr1611 packet bytes: %d", len(bytes)))
    }
    
    return packet, err
}

type uvr1611PacketDecoder struct {
    ByteConsumer
    consumer UVR1611PacketConsumer
    
    bytes []Byte
}

func NewUVR1611PacketDecoder(consumer UVR1611PacketConsumer) *uvr1611PacketDecoder {
    d := &uvr1611PacketDecoder{consumer: consumer}
    d.bytes = make([]Byte, 0, UVR1611PacketByteCount)
    
    return d
}

func (d *uvr1611PacketDecoder) Reset() {
    d.bytes = make([]Byte, 0, cap(d.bytes))
}

func (d *uvr1611PacketDecoder) Consume(b Byte) error {
    bytes := append(d.bytes, b)
    if len(bytes) == cap(d.bytes) {
        fmt.Printf("%v\n", bytes)
        d.Reset()
        packet, err := UVR1611PacketFromBytes(bytes)
        if err != nil {
            fmt.Println("Could not parse packet bytes", err)
        } else {
            d.consumer.Consume(packet)
        }
    } else {
        d.bytes = bytes
    }
    
    return nil
}