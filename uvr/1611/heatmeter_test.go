package uvr1611

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/brutella/gouvr/uvr"
)

func TestHeatmetersEnabled(t *testing.T) {
    b := uvr.Byte(0x2) // 0000 0010
    h1, h2 := AreHeatMetersEnabled(uvr.NewHeatMeterRegister(b))
    assert.False(t, h1)
    assert.True(t, h2)
}