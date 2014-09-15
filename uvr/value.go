package uvr

type Value struct {
    Low Byte
    High Byte
}

func NewValue(bytes []Byte) Value {
    return Value{
        Low: bytes[0],
        High: bytes[1],
    }
}

func Int16FromValue(value Value) int16 {
    return (int16(value.High << 8) | int16(value.Low))
}

type BigValue struct {
    Low Value
    High Value
}

func NewBigValue(bytes []Byte) BigValue {
    return BigValue{
        Low: NewValue(bytes[0:2]),
        High: NewValue(bytes[2:4]),
    }
}

func Int32FromBigValue(value BigValue) int32 {
    return (int32(value.High.High << 24) | int32(value.High.Low << 16) | int32(value.Low.High << 8) | int32(value.Low.Low))
}