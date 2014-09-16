package uvr

import(
    "fmt"
)

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
        currentPower_kW: NewBigValue(bytes[0:4]),
        power_kWh: NewValue(bytes[4:6]),
        power_MWh: NewValue(bytes[6:8]),
    }
}

func (h *HeatMeterValue) ToString() string {
    // TODO
    // Decode bytes correctly based on specification
    current_kW := Int32FromBigValue(h.currentPower_kW)/100
    power_kWh := Int16FromValue(h.power_kWh)/10
    power_MWh := Int16FromValue(h.power_MWh)/10
    
    return fmt.Sprintf("Current %d kW | %d kWh | %d MWh", current_kW, power_kWh, power_MWh)
}