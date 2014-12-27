package uvr1611

import (
	"fmt"
	"github.com/brutella/gouvr/uvr"
)

// IsUnusedInputValue returns true when the value for an input is unused
func IsUnusedInputValue(value uvr.Value) bool {
	return InputTypeFromValue(value) == InputTypeUnused
}

// RoomTemperatureModeFromValue returns the room temperature mode from the value.
func RoomTemperatureModeFromValue(value uvr.Value) RoomTemperatureMode {
	if InputTypeFromValue(value) == InputTypeRoomTemperature {
		return RoomTemperatureMode(value.High & InputTypeRoomTemperatureOperationModeMask)
	}

	return RoomTemperatureModeUndefined
}

// InputTypeFromValue returns the input type of the value.
func InputTypeFromValue(value uvr.Value) InputType {
	return InputType(value.High & InputTypeMask)
}

func IsDigitalInputValueOn(value uvr.Value) bool {
	input_type, f := DecodeInputValue(value)
	return input_type == InputTypeDigital && f == 1.0
}

// DecodeInputValue returns the input type and the float value.
func DecodeInputValue(value uvr.Value) (InputType, float32) {
	input_type := InputTypeFromValue(value)

	var result float32
	var low_byte = uvr.Byte(0)
	var high_byte = uvr.Byte(0)

	if input_type != InputTypeUnused {
		low_byte = value.Low & InputTypeLowValueMask
		// Room temperature input is handled differently
		if input_type == InputTypeRoomTemperature {
			high_byte = value.High & InputTypeRoomTemperatureHighValueMask
		} else {
			// Extract only interesting bits
			high_byte = value.High & InputTypeHighValueMask
		}

		if value.High&InputTypeSignMask == InputTypeSignMask { // 1000 0000
			if input_type == InputTypeDigital {
				// bit 7 = 1 (On)
				result = 1.0
			} else {
				// For negative values, set bit 4,5,6 in high byte to 1
				high_byte |= 0xF0 // 1111 0000
				result = float32(low_byte) + float32(high_byte)*256 - 65536
				result /= 10.0
			}
		} else {
			if input_type == InputTypeDigital {
				// bit 7 = 0 (Off)
				result = 0.0
			} else {
				// For positive values, set bit 4,5,6 in high byte to 0
				high_byte &= 0x0F // 0000 1111
				result = float32(low_byte) + float32(high_byte)*256
				result /= 10.0
			}
		}
	}

	return input_type, result
}

// InputValueToString returns the value properly formatted as string (e.g. 10.5 °C).
func InputValueToString(v uvr.Value) string {
	input_type, value := DecodeInputValue(v)

	var suffix string
	switch input_type {
	case InputTypeUnused:
		return "?"
	case InputTypeDigital:
		if value == 1.0 {
			return "On"
		}
		return "Off"
	case InputTypeRoomTemperature:
		suffix = "°C"
		suffix += " "
		switch RoomTemperatureModeFromValue(v) {
		case RoomTemperatureModeAutomatic:
			suffix += "[Auto]"
		case RoomTemperatureModeNormal:
			suffix += "[Normal]"
		case RoomTemperatureModeLowering:
			suffix += "[Lowering]"
		case RoomTemperatureModeStandby:
			suffix += "[Standby]"
		}
	case InputTypeTemperature:
		suffix = "°C"
	case InputTypeVolumeFlow:
		suffix = "l/h"
	case InputTypeRadiation:
		suffix = "W/m^2"
	}

	return fmt.Sprintf("%.01f %s", value, suffix)
}
