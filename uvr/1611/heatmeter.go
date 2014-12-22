package uvr1611

import (
	"github.com/brutella/gouvr/uvr"
)

// AreHeatMetersEnabled returns a boolean tuple for heatmeter 1 and 2 which are true
// if the corresponding heatmeter is used
func AreHeatMetersEnabled(r uvr.HeatMeterRegister) (bool, bool) {
	h1 := r.Value&0x1 == 0x1
	h2 := r.Value&0x2 == 0x2
	return h1, h2
}
