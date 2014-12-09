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
    assert.False(t, outlets[0].enabled)
    assert.False(t, outlets[1].enabled)
    assert.False(t, outlets[2].enabled)
    assert.True(t, outlets[3].enabled)
    assert.False(t, outlets[4].enabled)
    assert.True(t, outlets[5].enabled)
    assert.True(t, outlets[6].enabled)
    assert.False(t, outlets[7].enabled)
    assert.False(t, outlets[8].enabled)
    assert.False(t, outlets[9].enabled)
    assert.False(t, outlets[10].enabled)
    assert.False(t, outlets[11].enabled)
    assert.False(t, outlets[12].enabled)
}