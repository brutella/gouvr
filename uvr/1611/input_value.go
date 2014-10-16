package uvr1611

import(
    "github.com/brutella/gouvr/uvr"
    "fmt"
)

// Returns the input type of the value.
func InputTypeFromValue(value uvr.Value) InputType {
    return InputType(value.High & InputTypeMask)
}

// Returns the input type and the float value.
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

// Returns the value properly formatted as string (e.g. 10.5 °C).
func InputValueToString(v uvr.Value) string {
    input_type, value := DecodeInputValue(v)
    
    var suffix string
    switch input_type {
    case InputTypeUnused:
        return "?"
    case InputTypeDigital:
        suffix = ""
    case InputTypeRoomTemperature:
        fallthrough
    case InputTypeTemperature:
        suffix = "°C"
    case InputTypeVolumeFlow:
        suffix = "l/h"
    case InputTypeRadiation:
        suffix = "W/m^2"
    }
    
    return fmt.Sprintf("%.01f %s", value, suffix) 
}