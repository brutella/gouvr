package uvr1611

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/brutella/gouvr/uvr"
)

func TestOutlets(t *testing.T) {
    value := uvr.NewValue([]uvr.Byte{
        uvr.Byte(0x68), // 0110 1000
        uvr.Byte(0x00), // 0000 0000
    })
    
    outlets := OutletsFromValue(value)
    assert.Equal(t, len(outlets), 13)
    assert.False(t, outlets[0].Enabled)
    assert.False(t, outlets[1].Enabled)
    assert.False(t, outlets[2].Enabled)
    assert.True(t, outlets[3].Enabled)
    assert.False(t, outlets[4].Enabled)
    assert.True(t, outlets[5].Enabled)
    assert.True(t, outlets[6].Enabled)
    assert.False(t, outlets[7].Enabled)
    assert.False(t, outlets[8].Enabled)
    assert.False(t, outlets[9].Enabled)
    assert.False(t, outlets[10].Enabled)
    assert.False(t, outlets[11].Enabled)
    assert.False(t, outlets[12].Enabled)
}