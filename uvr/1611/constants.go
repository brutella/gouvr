package uvr1611

import (
	"github.com/brutella/gouvr/uvr"
)

const (
	// SignalFrequency is the frequency of the UVR1611 data bus
	SignalFrequency = 488.0
)

// Bit masks
const (
	InputTypeSignMask      = 0x80 // 1000 0000
	InputTypeMask          = 0x70 // 0111 0000
	InputTypeHighValueMask = 0x0F // 0000 1111
	InputTypeLowValueMask  = 0xFF // 1111 1111

	InputTypeRoomTemperatureHighValueMask     = 0x01 // 0000 0001
	InputTypeRoomTemperatureOperationModeMask = 0x06 // 0000 0110
)

// Input Types
type InputType uvr.Byte

const (
	InputTypeUnused          = InputType(0x00)
	InputTypeDigital         = InputType(0x10) // 0001 0000
	InputTypeTemperature     = InputType(0x20) // 0010 0000
	InputTypeVolumeFlow      = InputType(0x30) // 0011 0000
	InputTypeRadiation       = InputType(0x60) // 0110 0000
	InputTypeRoomTemperature = InputType(0x70) // 0111 0000
)

// Room Temperature Operation Modes
type RoomTemperatureMode uvr.Byte

const (
	RoomTemperatureModeUndefined = RoomTemperatureMode(0xFF)
	RoomTemperatureModeAutomatic = RoomTemperatureMode(0x00) // 0000 0000
	RoomTemperatureModeNormal    = RoomTemperatureMode(0x02) // 0000 0010
	RoomTemperatureModeLowering  = RoomTemperatureMode(0x04) // 0000 0100
	RoomTemperatureModeStandby   = RoomTemperatureMode(0x06) // 0000 0110
)
