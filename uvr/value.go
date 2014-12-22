package uvr

// A value consists of 2 bytes
type Value struct {
	Low  Byte
	High Byte
}

// NewValue returns a value from 2 bytes
func NewValue(bytes []Byte) Value {
	return Value{
		Low:  bytes[0],
		High: bytes[1],
	}
}

func Int16FromValue(value Value) int16 {
	return (int16(value.High)<<8 | int16(value.Low))
}

func UInt16FromValue(value Value) uint16 {
	return (uint16(value.High)<<8 | uint16(value.Low))
}

// A big value consists of 2 values (= 4 bytes)
type BigValue struct {
	Low  Value
	High Value
}

// NewBigValue returns a big value from 4 bytes
func NewBigValue(bytes []Byte) BigValue {
	return BigValue{
		Low:  NewValue(bytes[0:2]),
		High: NewValue(bytes[2:4]),
	}
}

func Int32FromBigValue(value BigValue) int32 {
	return (int32(value.High.High)<<24 | int32(value.High.Low)<<16 | int32(value.Low.High)<<8 | int32(value.Low.Low))
}

func UInt32FromBigValue(value BigValue) uint32 {
	return (uint32(value.High.High)<<24 | uint32(value.High.Low)<<16 | uint32(value.Low.High)<<8 | uint32(value.Low.Low))
}
