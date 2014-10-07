package uvr1611

import(
    "github.com/brutella/gouvr/uvr"
)

const (
    SignalFrequency = 488
)

// Masks
const(
    InputTypeMask                         = 0x70 // 0111 0000
    InputTypeHighValueMask                = 0x0F // 0000 1111
    InputTypeLowValueMask                 = 0xFF // 1111 1111
    
    InputTypeRoomTemperatureHighValueMask = 0x01 // 0000 0001
    InputTypeRoomTemperatureOperationModeMask = 0x06 // 0000 0110
)

// Input Type
const (
    InputTypeUnused          = 0x00
    InputTypeDigital         = 0x10 // 0001 0000
    InputTypeTemperature     = 0x20 // 0010 0000
    InputTypeVolumeFlow      = 0x30 // 0011 0000
    InputTypeRadiation       = 0x60 // 0110 0000
    InputTypeRoomTemperature = 0x70 // 0111 0000
)

// Room Temperature Operation Mode
const (
    RoomTemperatureModeAutomatic = 0x00 // 0000 0000
    RoomTemperatureModeNormal    = 0x02 // 0000 0010
    RoomTemperatureModeLowering  = 0x04 // 0000 0100
    RoomTemperatureModeStandby   = 0x06 // 0000 0110
)

type InputType uvr.Byte