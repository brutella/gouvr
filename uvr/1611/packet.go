package uvr1611

import (
    "gouvr/uvr"
    "fmt"
    "errors"
)

const PacketByteCount = 64

type Packet struct {
    deviceType uvr.DeviceType           // 1
    deviceTypeInverted uvr.DeviceType   // 2
    unkown uvr.Byte                     // 3
    timestamp uvr.Timestamp             // 4, 5, 6, 7, 8
    input1 uvr.Value        // 9, 10  (Aussentemperatur)
    input2 uvr.Value        // 11, 12 (Fussbodenheizung Vorlauf)
    input3 uvr.Value        // 13, 14 (Buffer Warmwasser - Oben)
    input4 uvr.Value        // 15, 16 (Buffer Warmwasser - Mitte)
    input5 uvr.Value        // 17, 18 (Buffer Warmwasser - Unten)
    input6 uvr.Value        // 19, 20 (Raumtemperatur)
    input7 uvr.Value        // 21, 22 (Wärmetauscher Sekundär)
    input8 uvr.Value        // 23, 24
    input9 uvr.Value        // 25, 26
    input10 uvr.Value       // 27, 28
    input11 uvr.Value       // 29, 30
    input12 uvr.Value       // 31, 32
    input13 uvr.Value       // 33, 34
    input14 uvr.Value       // 35, 36
    input15 uvr.Value       // 37, 38
    input16 uvr.Value       // 39, 40
    outgoing uvr.Value        // 41, 42
    speedStepA1 uvr.Byte     // 43
    speedStepA2 uvr.Byte     // 44
    speedStepA6 uvr.Byte     // 45
    speedStepA7 uvr.Byte     // 46
    heatRegister uvr.HeatMeterRegister // 47
    heatMeter1 uvr.HeatMeterValue // 48, 49, 50, 51 | 52, 53 | 54, 55
    heatMeter2 uvr.HeatMeterValue // 56, 57, 58, 59 | 60, 61 | 62, 63
    checksum uvr.Byte       // 64
}

func (p *Packet) Log() {
    fmt.Println("Device Type:", p.deviceType)
    fmt.Println("Inverted Device Type:", p.deviceTypeInverted)
    fmt.Println("Time:", p.timestamp.ToString())
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
    fmt.Println("Out:", uvr.Int16FromValue(p.outgoing))
    fmt.Println("A1:", uvr.NewSpeedStep(p.speedStepA1).ToString())
    fmt.Println("A2:", uvr.NewSpeedStep(p.speedStepA2).ToString())
    fmt.Println("A6:", uvr.NewSpeedStep(p.speedStepA6).ToString())
    fmt.Println("A7:", uvr.NewSpeedStep(p.speedStepA7).ToString())
    fmt.Println("Heat Register:", p.heatRegister)
    fmt.Println("Heat Meter 1:", p.heatMeter1.ToString())
    fmt.Println("Heat Meter 2:", p.heatMeter2.ToString())
}

type PacketConsumer interface {
    Consume(p Packet) error
}

// Returns the sum of all bytes modulo 256
func ChecksumFromBytes(bytes []uvr.Byte) uvr.Byte {
    var checksum uvr.Byte
    
    for _, b := range bytes {
        checksum += b
    }
    
    return checksum
}

// Returns a uvr1611 packet from a list of bytes or an error
func PacketFromBytes(bytes []uvr.Byte) (Packet, error){
    var err error
    var packet = Packet{}
    
    if len(bytes) == PacketByteCount {
        packet.deviceType = uvr.DeviceType(bytes[0])
        packet.deviceTypeInverted = uvr.DeviceType(bytes[1])
        packet.unkown = bytes[2]
        packet.timestamp = uvr.NewTimestamp(bytes[3:8])
        packet.input1 = uvr.NewValue(bytes[8:10])
        packet.input2 = uvr.NewValue(bytes[10:12])
        packet.input3 = uvr.NewValue(bytes[12:14])
        packet.input4 = uvr.NewValue(bytes[14:16])
        packet.input5 = uvr.NewValue(bytes[16:18])
        packet.input6 = uvr.NewValue(bytes[18:20])
        packet.input7 = uvr.NewValue(bytes[20:22])
        packet.input8 = uvr.NewValue(bytes[22:24])
        packet.input9 = uvr.NewValue(bytes[24:26])
        packet.input10 = uvr.NewValue(bytes[26:28])
        packet.input11 = uvr.NewValue(bytes[28:30])
        packet.input12 = uvr.NewValue(bytes[30:32])
        packet.input13 = uvr.NewValue(bytes[32:34])
        packet.input14 = uvr.NewValue(bytes[34:36])
        packet.input15 = uvr.NewValue(bytes[36:38])
        packet.input16 = uvr.NewValue(bytes[38:40])
        packet.outgoing = uvr.NewValue(bytes[40:42])
        packet.speedStepA1 = bytes[42]
        packet.speedStepA2 = bytes[43]
        packet.speedStepA6 = bytes[44]
        packet.speedStepA7 = bytes[45]
        packet.heatRegister = uvr.NewHeatMeterRegister(bytes[46])
        packet.heatMeter1 = uvr.NewHeatMeterValue(bytes[47:55])
        packet.heatMeter2 = uvr.NewHeatMeterValue(bytes[55:63])
        packet.checksum = bytes[63]
        
        checksum := ChecksumFromBytes(bytes[0:63])
        if packet.checksum != checksum {
            err = errors.New(fmt.Sprintf("Invalid checksum %d should be %d.", checksum, packet.checksum))
        }
    } else {
        err = errors.New(fmt.Sprintf("Invalid number of uvr1611 packet bytes: %d.", len(bytes)))
    }
    
    return packet, err
}

type packetDecoder struct {
    uvr.ByteConsumer
    consumer PacketConsumer
    
    bytes []uvr.Byte
}

func NewPacketDecoder(consumer PacketConsumer) *packetDecoder {
    d := &packetDecoder{consumer: consumer}
    d.bytes = make([]uvr.Byte, 0, PacketByteCount)
    
    return d
}

func (d *packetDecoder) Reset() {
    d.bytes = make([]uvr.Byte, 0, cap(d.bytes))
}

func (d *packetDecoder) Consume(b uvr.Byte) error {
    bytes := append(d.bytes, b)
    if len(bytes) == cap(d.bytes) {
        // fmt.Printf("%v\n", bytes)
        d.Reset()
        packet, err := PacketFromBytes(bytes)
        if err != nil {
            fmt.Println("[PACKET] Could not parse packet bytes.", err)
        } else {
            d.consumer.Consume(packet)
        }
    } else {
        d.bytes = bytes
    }
    
    return nil
}