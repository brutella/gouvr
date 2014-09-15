package uvr

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

func Int16FromValue(value Value) int16 {
    return (int16(value.high << 8) | int16(value.low))
}

func InputTypeFromValue(value Value) InputType {
    return InputType(value.high & InputTypeMask)
}

func DecodeInputValue(value Value) (InputType, float32) {
    input_type := InputTypeFromValue(value)
    
    var result float32
    var low_byte Byte = Byte(0)
    var high_byte Byte = Byte(0)
    if input_type != InputTypeUnused {
        low_byte = value.low & InputTypeLowValueMask
        // Room temperature input is handled differenctly
        if input_type == InputTypeRoomTemperature {
            high_byte = value.high & InputTypeRoomTemperatureHighValueMask
        } else {
            // Extract only interesting bits
            high_byte = value.high & InputTypeHighValueMask
        }
        
        if high_byte & 0x80 == 0x80 { // 1000 0000
            // For negative values, set bit 4,5,6 in high byte to 1
            high_byte |= 0x70 // 0111 0000
            result = float32(low_byte) + float32(high_byte) * 256 - 65536
        } else {
            // For positive values, set bit 4,5,6 in high byte to 0
            high_byte &= 0x0F // 0000 1111
            result = float32(low_byte) + float32(high_byte) * 256
        }
    }
    
    return input_type, result/10.0
}

type BigValue struct {
    low Value
    high Value
}

func NewBigValue(bytes []Byte) BigValue {
    return BigValue{
        low: NewValue(bytes[0:2]),
        high: NewValue(bytes[2:4]),
    }
}

func Int32FromBigValue(value BigValue) int32 {
    return (int32(value.high.high << 24) | int32(value.high.low << 16) | int32(value.low.high << 8) | int32(value.low.low))
}