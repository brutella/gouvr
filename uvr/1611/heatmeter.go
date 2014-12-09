package uvr1611

import(
    "github.com/brutella/gouvr/uvr"
)

// AreHeatMetersEnabled returns a bool tuple for heatmeter 1 and 2
func AreHeatMetersEnabled(r uvr.HeatMeterRegister) (bool, bool) {
    h1 := r.Value & 0x1 == 0x1
    h2 := r.Value & 0x2 == 0x2
    return h1, h2
}