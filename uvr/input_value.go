package uvr

import(
    "fmt"
)

func InputValueToString(v Value) string {
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