package uvr

import(
    "fmt"
)

type HeatMeterRegister struct {
    Value Byte
}

func NewHeatMeterRegister(value Byte) HeatMeterRegister {
    return HeatMeterRegister{ Value: value }
}

type HeatMeterValue struct {
    currentPower_kW BigValue
    power_kWh Value
    power_MWh Value
}

func NewHeatMeterValue(bytes []Byte) HeatMeterValue {
    return HeatMeterValue{
        currentPower_kW: NewBigValue(bytes[0:4]),
        power_kWh: NewValue(bytes[4:6]),
        power_MWh: NewValue(bytes[6:8]),
    }
}

func (h *HeatMeterValue) CurrentPower() float32 {
    value  := h.currentPower_kW
    i := int32(value.High.High) << 16 | int32(value.High.Low) << 8 | int32(value.Low.High)
    f := float32(i) + float32(value.Low.Low) * 10.0/256.0
    return f/100
}

func (h *HeatMeterValue) ToString() string {
    // TODO
    // Decode bytes correctly based on specification
    current_kW := h.CurrentPower()
    power_kWh := float32(Int16FromValue(h.power_kWh))/10.0
    power_MWh := Int16FromValue(h.power_MWh)
    
    return fmt.Sprintf("Current %f kW | %f kWh | %d MWh", current_kW, power_kWh, power_MWh)
}