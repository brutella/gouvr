package uvr

import (
    "math/big"
    "time"
    "fmt"
)

type Bit struct {
    Raw big.Word
    Timestamp time.Time
}

func NewBitFromWord(value big.Word) Bit {
    return Bit{Raw: value, Timestamp: time.Now()}
}

func NewBitFromInt(value int) Bit {
    return NewBitFromWord(big.Word(value))
}

type Byte uint8

type Timestamp struct {
    minute  Byte
    hour    Byte
    day     Byte
    month   Byte
    year    Byte // since 2000, e.g. 3 == 2003
}

func NewTimestamp(bytes []Byte) Timestamp {
    return Timestamp{
        minute: bytes[0],
        hour: bytes[1],
        day: bytes[2],
        month: bytes[3],
        year: bytes[4],
    }
}

func (t Timestamp) toString() string {
    return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:00", 2000 + int64(t.year), t.month, t.day, t.hour, t.minute)
}

type Value struct {
    low Byte
    high Byte
}

func NewValue(bytes []Byte) Value {
    return Value{
        low: bytes[0],
        high: bytes[1],
    }
}

type BigValue struct {
    low Value
    high Value
}

func NewBigValue(bytes []Byte) BigValue {
    return BigValue{
        low: NewValue(bytes[0:1]),
        high: NewValue(bytes[2:3]),
    }
}

type HeatMeterRegister struct {
    value Byte
}

func NewHeatMeterRegister(value Byte) HeatMeterRegister {
    return HeatMeterRegister{ value: value }
}

type HeatMeterValue struct {
    currentPower_kW BigValue // 0, 1, 2, 3
    power_kWh Value // 4, 5
    power_MWh Value // 6, 7
}

func NewHeatMeterValue(bytes []Byte) HeatMeterValue {
    return HeatMeterValue{
        currentPower_kW: NewBigValue(bytes[0:3]),
        power_kWh: NewValue(bytes[4:5]),
        power_MWh: NewValue(bytes[6:7]),
    }
}

type UVR1611Packet struct {
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
}
